package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"pengin_party/config"
	"pengin_party/di"
	"pengin_party/internal/infrastructure/redis"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	config.LoadConfig()
	// fmt.Println(config.DB.DNS())
	fmt.Println("APP_ENVは：" + os.Getenv("APP_ENV"))

	if err != nil {
		panic(err)
	}
	fmt.Println("DB_DATABASEは：" + os.Getenv("DB_DATABASE"))

	con, err := di.InitializeControllers()
	if err != nil {
		panic("DIの初期化に失敗しました。")
	}
	scon := con.ServerController

	// Redisクライアントの初期化
	ctx := context.Background()
	rdb, err := redis.NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}
	defer rdb.Close()
	defer con.DB.Close()

	router := gin.Default()

	// CORS設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// router.POST("/login", )
	router.POST("/users", scon.UserController.Create)      // ユーザーの登録（to MySQL）
	router.GET("/users/:uid", scon.UserController.IsExist) // ユーザーの存在確認
	// router.PUT("/rooms/{roomId}", )	// マッチング部屋を作成（to Redis）

	router.Run(":4000")
}
