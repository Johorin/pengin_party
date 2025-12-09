package controllers

import (
	"net/http"
	"pengin_party/internal/application/usecases/user/usecase"

	"github.com/gin-gonic/gin"
	// "github.com/redis/go-redis/v9"
)

type UserController interface {
	Create(c *gin.Context)
}

type userController struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserController(
	createUserUseCase *usecase.CreateUserUseCase,
) UserController {
	return &userController{
		createUserUseCase,
	}
}

func (uc *userController) Create(c *gin.Context) {
	var req struct {
		UserID   string `json:"user_id"`
		UserUUID string `json:"user_uuid"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserApiResponse{
		Data: UserResponse{
			ID: 1,
			UID: 11,
			UUID: 111,
			Name: "yuya",
		},
	})
}
