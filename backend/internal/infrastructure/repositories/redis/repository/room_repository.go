package repository

import (
	"context"
	"fmt"
	"pengin_party/internal/domain/room"
	"pengin_party/internal/infrastructure/repositories/redis"
)

type RoomRepository struct {
	redisClient redis.RedisInterface
}

func NewRoomRepository(
	redisClient redis.RedisInterface,
) room.RoomRepository {
	return &RoomRepository{redisClient}
}

func (rRepo *RoomRepository) CreateRoom(ctx context.Context, roomId *string, userUid string) (*string, error) {
	// ホストのセット
	// room:123:host = user_1
	err := rRepo.redisClient.GetRedis().Set(ctx, "room:" + *roomId + ":host", userUid, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("ホストのセットに失敗しました: %w", err)
	}

	// 部屋の作成
	// room:123:users = {user_1, user_2, ...}
	err = rRepo.redisClient.GetRedis().RPush(ctx, "room:" + *roomId + ":users", userUid).Err()
	if err != nil {
		return nil, fmt.Errorf("部屋の作成に失敗しました: %w", err)
	}

	return roomId, nil
}