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
	IsExist(c *gin.Context)
}

type userController struct {
	createUserUseCase *usecase.CreateUserUseCase
	isExistUserUseCase *usecase.IsExistUserUseCase
}

func NewUserController(
	createUserUseCase *usecase.CreateUserUseCase,
	isExistUserUseCase *usecase.IsExistUserUseCase,
) UserController {
	return &userController{
		createUserUseCase,
		isExistUserUseCase,
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

func (uc *userController) IsExist(c *gin.Context) {
	// var req types.LoginUser
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	panic("リクエストのフォーマットが違います")
	// }
	uid := c.Param("uid")

	isExist, err := uc.isExistUserUseCase.Execute(c.Request.Context(), uid)
	if err != nil {
		panic("")
	}

	c.JSON(http.StatusOK, SearchUserApiResponse{
		Data: SearchUserResponse{
			IsExist: isExist,
		},
	})
}
