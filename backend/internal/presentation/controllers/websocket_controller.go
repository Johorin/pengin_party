package controllers

import (
	"log"
	"pengin_party/internal/application/usecases/websocket/usecase"

	"github.com/gin-gonic/gin"
)

type WebSocketController interface {
	Handle(c *gin.Context)
}

type webSocketController struct {
	handleWebSocketUseCase *usecase.HandleWebSocketUseCase
}

func NewWebSocketController(
	handleWebSocketUseCase *usecase.HandleWebSocketUseCase,
) WebSocketController {
	return &webSocketController{
		handleWebSocketUseCase: handleWebSocketUseCase,
	}
}

func (wsc *webSocketController) Handle(c *gin.Context) {
	writer  := c.Writer
	request := c.Request
	err := wsc.handleWebSocketUseCase.Execute(writer, request)
	if err != nil {
		log.Printf("WebSocket接続エラー: %v", err)
		return
	}
}