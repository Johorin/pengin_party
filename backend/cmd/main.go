package main

import (
    "context"
	"github.com/gin-gonic/gin"
	"pengin_party/internal/infrastructure/redis"
)

func main() {
	ctx := context.Background()
	rdb, err :=  redis.NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}
	defer rdb.Close()

	router := gin.Default()
	// router.POST("/login", )
	router.POST("/users", )			// ユーザーの登録（to MySQL）
	router.PUT("/rooms/{roomId}", )	// マッチング部屋を作成（to Redis）
}