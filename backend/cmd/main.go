package main

import (
	"context"
	"fmt"
	"os"
	"pengin_party/config"
	"pengin_party/internal/application/usecases/user/usecase"
	"pengin_party/internal/infrastructure/redis"
	"pengin_party/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "pengin_party/di"
)

func main() {
	config.LoadConfig()
	// fmt.Println(config.DB.DNS())
	fmt.Println("APP_ENVは：" + os.Getenv("APP_ENV"))

	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB_DATABASEは：" + os.Getenv("DB_DATABASE"))

	// Redisクライアントの初期化
	ctx := context.Background()
	rdb, err :=  redis.NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}
	defer rdb.Close()

	user_uc   := usecase.NewCreateUserUseCase()
	user_con  := controllers.NewUserController(user_uc)
	// user_repo := 

	router := gin.Default()
	// router.POST("/login", )
	router.POST("/users", user_con.Create)			// ユーザーの登録（to MySQL）
	// router.PUT("/rooms/{roomId}", )	// マッチング部屋を作成（to Redis）
}