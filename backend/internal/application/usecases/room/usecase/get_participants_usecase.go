package usecase

import (
	"context"
	"pengin_party/internal/domain/room"

	"github.com/cockroachdb/errors"
)

type GetParticipantsUseCase struct {
	roomRepo room.RoomRepository
}

func NewGetParticipantsUseCase(
	roomRepo room.RoomRepository,
) *GetParticipantsUseCase {
	return &GetParticipantsUseCase{
		roomRepo: roomRepo,
	}
}

func (uc *GetParticipantsUseCase) Execute(
	ctx context.Context,
	roomId string,
) ([]string, error) {
	participants, err := uc.roomRepo.GetParticipants(ctx, roomId)
	if err != nil {
		return nil, errors.Wrap(err, "参加者リストの取得に失敗しました")
	}

	return participants, nil
}
