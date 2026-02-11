package websocket

import "github.com/gorilla/websocket"

type ConnectionRepository interface {
	AddConnection(roomId string, conn *websocket.Conn)
	RemoveConnection(roomId string, conn *websocket.Conn)
	GetConnections(roomId string) []*websocket.Conn
	RemoveConnectionByConn(conn *websocket.Conn)
}