package usecase

import (
	"context"
	"pengin_party/internal/domain/room"
	"pengin_party/internal/infrastructure/repositories/redis"

	"github.com/cockroachdb/errors"
)

type CreateRoomUseCase struct {
	redisClient redis.RedisInterface
	roomRepo    room.RoomRepository
}

func NewCreateRoomUseCase(
	redisClient redis.RedisInterface,
	roomRepo room.RoomRepository,
) *CreateRoomUseCase {
	return &CreateRoomUseCase{
		redisClient: redisClient,
		roomRepo:    roomRepo,
	}
}

func (uc *CreateRoomUseCase) Execute(
	ctx context.Context,
	roomId *string,
) (*string, error) {
	roomId, err := uc.roomRepo.CreateRoom(ctx, roomId)
	if err != nil {
		return nil, errors.Wrap(err, "部屋の作成に失敗しました")
	}

	return roomId, nil
}
