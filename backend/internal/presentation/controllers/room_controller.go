package controllers

import (
	"net/http"
	"pengin_party/internal/application/usecases/room/usecase"

	"github.com/gin-gonic/gin"
)

type RoomController interface {
	Create(c *gin.Context)
}

type roomController struct {
	createRoomUseCase *usecase.CreateRoomUseCase
}

func NewRoomController(
	createRoomUseCase *usecase.CreateRoomUseCase,
) RoomController {
	return &roomController{
		createRoomUseCase,
	}
}

func (r *roomController) Create(c *gin.Context) {
	var req struct {
		RoomID *string `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		panic("リクエストボディの取得に失敗しました。")
	}

	roomId, err := r.createRoomUseCase.Execute(c.Request.Context(), req.RoomID)
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