package controllers

import (
	"net/http"
	"pengin_party/internal/application/usecases/user/usecase"

	"github.com/gin-gonic/gin"
	// "github.com/redis/go-redis/v9"
	"pengin_party/pkg/types"
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
	var req types.LoginUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := uc.createUserUseCase.Execute(c.Request.Context(), req.Name, req.Email, req.UID, req.UUID)
	if err != nil {
		panic("ユーザー作成ユースケースの実行に失敗しました。")
	}

	c.JSON(http.StatusOK, CreateUserApiResponse{
		Data: CreateUserResponse{
			ID: userId,
		},
	})
}
