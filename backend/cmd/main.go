package main

import (
    "context"
	"github.com/gin-gonic/gin"
	"pengin_party/internal/infrastructure/redis"
)

func main() {
	ctx := context.Background()
	rds, err :=  redis.NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}
	defer rds.Close()

	router := gin.Default()
	router.POST("/login", )
}