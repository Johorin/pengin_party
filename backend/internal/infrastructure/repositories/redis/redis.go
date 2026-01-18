package redis

import (
    "context"
    "fmt"
	"time"
    "github.com/redis/go-redis/v9"
)

type RedisInterface interface {
    GetRedis() *redis.Client
    Close()    error
}

type RedisClient struct {
    client *redis.Client
    ctx    context.Context
}

func Init() (RedisInterface, error) {
    ctx := context.Background()

    // クライアントの初期化
    client := redis.NewClient(&redis.Options{
        Addr:         "redis:6379",
        Password:     "",
        DB:           0,
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolSize:     10,
        PoolTimeout:  30 * time.Second,
    })

    // 接続テスト
    if err := client.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("failed to connect to redis: %v", err)
    }

    return &RedisClient{
        client: client,
        // ctx:    ctx,
    }, nil
}

func (rc *RedisClient) GetRedis() *redis.Client {
    return rc.client
}

func (rc *RedisClient) Close() error {
    return rc.client.Close()
}