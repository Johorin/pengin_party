package services

import (
	"encoding/json"
	"log"
	wsDomain "pengin_party/internal/domain/websocket"

	"github.com/gorilla/websocket"
)

type WebSocketService interface {
	BroadcastToRoom(roomId string, message interface{}) error
}

type webSocketService struct {
	connRepo wsDomain.ConnectionRepository
}

func NewWebSocketService(connRepo wsDomain.ConnectionRepository) WebSocketService {
	return &webSocketService{connRepo: connRepo}
}

func (s *webSocketService) BroadcastToRoom(roomId string, message interface{}) error {
	conns := s.connRepo.GetConnections(roomId)
	if len(conns) == 0 {
		return nil
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Printf("WebSocket送信エラー: %v", err)
			conn.Close()
			s.connRepo.RemoveConnectionByConn(conn)
		}
	}
	return nil
}
