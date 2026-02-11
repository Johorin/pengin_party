package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"pengin_party/internal/application/services"
	"pengin_party/internal/application/usecases/room/usecase"
	"pengin_party/internal/infrastructure/firebase"
	"pengin_party/internal/infrastructure/middleware"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoomController interface {
	Create(c *gin.Context)
	Join(c *gin.Context)
}

type roomController struct {
	createRoomUseCase      *usecase.CreateRoomUseCase
	joinRoomUseCase        *usecase.JoinRoomUseCase
	getParticipantsUseCase *usecase.GetParticipantsUseCase
	firebase               firebase.FirebaseInterface
	websocketService       services.WebSocketService
}

func NewRoomController(
	createRoomUseCase *usecase.CreateRoomUseCase,
	joinRoomUseCase *usecase.JoinRoomUseCase,
	getParticipantsUseCase *usecase.GetParticipantsUseCase,
	firebase firebase.FirebaseInterface,
	websocketService services.WebSocketService,
) RoomController {
	return &roomController{
		createRoomUseCase:      createRoomUseCase,
		joinRoomUseCase:        joinRoomUseCase,
		getParticipantsUseCase: getParticipantsUseCase,
		firebase:               firebase,
		websocketService:       websocketService,
	}
}

func (r *roomController) Create(c *gin.Context) {
	var req struct {
		RoomID *string `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		panic("リクエストボディの取得に失敗しました。")
	}

	roomId := req.RoomID
	userUid := c.GetString(middleware.UserUID)
	if userUid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}

	roomId, err := r.createRoomUseCase.Execute(c.Request.Context(), roomId, userUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "部屋作成ユースケースの実行に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, CreateRoomApiResponse{
		Data: CreateRoomResponse{
			RoomID: roomId,
		},
	})
}

func (r *roomController) Join(c *gin.Context) {
	var req struct {
		RoomID string `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストボディの取得に失敗しました"})
		return
	}

	roomId := req.RoomID
	if roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "部屋IDが指定されていません"})
		return
	}

	userUid := c.GetString(middleware.UserUID)
	if userUid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}

	err := r.joinRoomUseCase.Execute(c.Request.Context(), roomId, userUid)
	if err != nil {
		// 部屋が見つからない場合のエラーハンドリング
		errMsg := err.Error()
		if strings.Contains(errMsg, "部屋が見つかりませんでした") {
			c.JSON(http.StatusNotFound, gin.H{"error": "部屋が見つかりませんでした"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "部屋への参加に失敗しました"})
		return
	}

	// WebSocketサーバーに参加者追加を通知（参加者リストも含める）
	go func() {
		websocketURL := os.Getenv("WEBSOCKET_URL")
		if websocketURL == "" {
			websocketURL = "http://localhost:3001"
		}

		// 参加者リストを取得
		ctx := context.Background()
		userUids, err := r.getParticipantsUseCase.Execute(ctx, roomId)
		if err != nil {
			// 参加者リストの取得に失敗しても通知は送信（エラーログのみ）
			return
		}

		// Firebase Admin SDKを使って各UIDからユーザー名を取得
		authClient, err := r.firebase.GetFirebase().Auth(ctx)
		if err != nil {
			return
		}

		participantNames := make([]string, 0, len(userUids))
		for _, uid := range userUids {
			userRecord, err := authClient.GetUser(ctx, uid)
			if err != nil {
				participantNames = append(participantNames, uid)
				continue
			}

			name := userRecord.DisplayName
			if name == "" {
				name = userRecord.Email
			}
			if name == "" {
				name = uid
			}
			participantNames = append(participantNames, name)
		}

		notificationData := map[string]interface{}{
			"roomId":       roomId,
			"participants": participantNames,
		}
		jsonData, _ := json.Marshal(notificationData)

		req, _ := http.NewRequest("POST", websocketURL+"/notify-participant-joined", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			// WebSocketサーバーへの通知が失敗してもエラーにしない（ログのみ）
		}
	}()

	go func() {
		message := map[string]interface{}{
			"type": "participants-updated",
			"roomId": roomId,
			"participants": "test",
		}
		r.websocketService.BroadcastToRoom(roomId, message)
	}()

	c.JSON(http.StatusOK, JoinRoomApiResponse{
		Data: JoinRoomResponse{
			RoomID: roomId,
		},
	})
}
