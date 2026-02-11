package usecase

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	wsDomain "pengin_party/internal/domain/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 本番環境では適切に設定
	},
}

type HandleWebSocketUseCase struct {
	connRepo wsDomain.ConnectionRepository
}

func NewHandleWebSocketUseCase(
	connRepo wsDomain.ConnectionRepository,
) *HandleWebSocketUseCase {
	return &HandleWebSocketUseCase{
		connRepo: connRepo,
	}
}

func (uc *HandleWebSocketUseCase) Execute(
	w http.ResponseWriter,
	r *http.Request,
) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	roomId := r.URL.Query().Get("roomId")
	if roomId == "" {
		return nil
	}

	uc.connRepo.AddConnection(roomId, conn)
	defer func() {
		uc.connRepo.RemoveConnection(roomId, conn)
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ReadMessage Error: %v", err)
			break
		}
	}
	return nil
}