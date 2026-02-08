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
	err := rRepo.redisClient.GetRedis().Set(ctx, "room:"+*roomId+":host", userUid, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("ホストのセットに失敗しました: %w", err)
	}

	// 部屋の作成
	// room:123:users = {user_1, user_2, ...}
	err = rRepo.redisClient.GetRedis().RPush(ctx, "room:"+*roomId+":users", userUid).Err()
	if err != nil {
		return nil, fmt.Errorf("部屋の作成に失敗しました: %w", err)
	}

	return roomId, nil
}

func (rRepo *RoomRepository) JoinRoom(ctx context.Context, roomId string, userUid string) error {
	// 部屋が存在するかチェック（room:{roomId}:hostが存在するか）
	hostKey := "room:" + roomId + ":host"
	exists, err := rRepo.redisClient.GetRedis().Exists(ctx, hostKey).Result()
	if err != nil {
		return fmt.Errorf("部屋の存在確認に失敗しました: %w", err)
	}
	if exists == 0 {
		return fmt.Errorf("部屋が見つかりませんでした")
	}

	// 部屋のユーザーリストに追加（右からプッシュ）
	usersKey := "room:" + roomId + ":users"
	err = rRepo.redisClient.GetRedis().RPush(ctx, usersKey, userUid).Err()
	if err != nil {
		return fmt.Errorf("部屋への参加に失敗しました: %w", err)
	}

	return nil
}

func (rRepo *RoomRepository) GetParticipants(ctx context.Context, roomId string) ([]string, error) {
	// 部屋が存在するかチェック
	hostKey := "room:" + roomId + ":host"
	exists, err := rRepo.redisClient.GetRedis().Exists(ctx, hostKey).Result()
	if err != nil {
		return nil, fmt.Errorf("部屋の存在確認に失敗しました: %w", err)
	}
	if exists == 0 {
		return nil, fmt.Errorf("部屋が見つかりませんでした")
	}

	// 部屋のユーザーリストを取得
	usersKey := "room:" + roomId + ":users"
	userUids, err := rRepo.redisClient.GetRedis().LRange(ctx, usersKey, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("参加者リストの取得に失敗しました: %w", err)
	}

	return userUids, nil
}
