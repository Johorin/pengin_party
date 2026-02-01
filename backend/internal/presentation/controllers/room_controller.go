package controllers

import (
	"net/http"
	"pengin_party/internal/application/usecases/room/usecase"
	"pengin_party/internal/infrastructure/middleware"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoomController interface {
	Create(c *gin.Context)
	Join(c *gin.Context)
}

type roomController struct {
	createRoomUseCase *usecase.CreateRoomUseCase
	joinRoomUseCase   *usecase.JoinRoomUseCase
}

func NewRoomController(
	createRoomUseCase *usecase.CreateRoomUseCase,
	joinRoomUseCase *usecase.JoinRoomUseCase,
) RoomController {
	return &roomController{
		createRoomUseCase: createRoomUseCase,
		joinRoomUseCase:   joinRoomUseCase,
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

	c.JSON(http.StatusOK, JoinRoomApiResponse{
		Data: JoinRoomResponse{
			RoomID: roomId,
		},
	})
}
