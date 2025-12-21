package controllers

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// func RoomController(c *gin.Context, rdb *RedisClient) {
// 	var req struct {
// 		UserID string   `json:"user_id`
// 		UserUUID string `json:"user_uuid"`
// 	}
// 	roomId := c.Param("roomId")
// 	c.ShouldBindJSON(&req)

// 	// Redisに部屋番号を保存

// 	c.JSON(http.StatusOK, ApiResponse{
// 		Data: RoomResponse{
// 			RoomId: roomId,
// 			Status: "active",
// 		},
// 	})
// }