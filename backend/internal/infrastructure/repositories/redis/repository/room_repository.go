package repository

import (
	"context"
	"fmt"
	"pengin_party/internal/infrastructure/repositories/redis"
	"pengin_party/internal/domain/room"
)

type RoomRepository struct {
	redisClient redis.RedisInterface
}

func NewRoomRepository(
	redisClient redis.RedisInterface,
) room.RoomRepository {
	return &RoomRepository{redisClient}
}

func (rRepo *RoomRepository) CreateRoom(ctx context.Context, roomId *string) (*string, error) {
	// ホストのセット
	// room:123:host = user_1
	err := rRepo.redisClient.GetRedis().Set(ctx, "room:123:host", "user1uid", 0).Err()
	if err != nil {
		return nil, fmt.Errorf("ホストのセットに失敗しました: %w", err)
	}

	// 部屋の作成
	// room:123:users = {user_1, user_2, ...}
	err = rRepo.redisClient.GetRedis().RPush(ctx, "room:123:users", "user1uid", "user2uid").Err()
	if err != nil {
		return nil, fmt.Errorf("部屋の作成に失敗しました: %w", err)
	}

	return roomId, nil
}