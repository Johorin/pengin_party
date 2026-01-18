package room

import "context"

type RoomRepository interface {
	CreateRoom(ctx context.Context, roomId *string) (*string, error)
}
