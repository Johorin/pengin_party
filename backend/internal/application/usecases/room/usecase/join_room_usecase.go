package usecase

import (
	"context"
	"pengin_party/internal/domain/room"

	"github.com/cockroachdb/errors"
)

type JoinRoomUseCase struct {
	roomRepo room.RoomRepository
}

func NewJoinRoomUseCase(
	roomRepo room.RoomRepository,
) *JoinRoomUseCase {
	return &JoinRoomUseCase{
		roomRepo: roomRepo,
	}
}

func (uc *JoinRoomUseCase) Execute(
	ctx context.Context,
	roomId string,
	userUid string,
) error {
	err := uc.roomRepo.JoinRoom(ctx, roomId, userUid)
	if err != nil {
		return errors.Wrap(err, "部屋への参加に失敗しました")
	}

	return nil
}
