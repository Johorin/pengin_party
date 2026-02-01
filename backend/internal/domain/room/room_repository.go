package room

import "context"

type RoomRepository interface {
	CreateRoom(ctx context.Context, roomId *string, userUid string) (*string, error)
	JoinRoom(ctx context.Context, roomId string, userUid string) error
}
