package websocket

import (
	wsDomain "pengin_party/internal/domain/websocket"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionRepository struct {
	connections map[string][]*websocket.Conn
	connToRoom  map[*websocket.Conn]string
	mu          sync.RWMutex
}

func NewConnectionRepository() wsDomain.ConnectionRepository {
	return &connectionRepository{
		connections: make(map[string][]*websocket.Conn),
		connToRoom:  make(map[*websocket.Conn]string),
	}
}

func (r *connectionRepository) AddConnection(roomId string, conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.connections[roomId] = append(r.connections[roomId], conn)
	r.connToRoom[conn] = roomId
}

func (r *connectionRepository) RemoveConnection(roomId string, conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	conns := r.connections[roomId]
	for i, c := range conns {
		if c == conn {
			r.connections[roomId] = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	if len(r.connections[roomId]) == 0 {
		delete(r.connections, roomId)
	}
	delete(r.connToRoom, conn)
}

func (r *connectionRepository) GetConnections(roomId string) []*websocket.Conn {
	r.mu.RLock()
	defer r.mu.RUnlock()
	conns := r.connections[roomId]
	result := make([]*websocket.Conn, len(conns))
	copy(result, conns)
	return result
}

func (r *connectionRepository) RemoveConnectionByConn(conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	roomId, ok := r.connToRoom[conn]
	if !ok {
		return
	}
	r.RemoveConnection(roomId, conn)
}
