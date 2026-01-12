package controllers

type ServerController struct {
	UserController UserController
}

func NewServerController(
	userController UserController,
) *ServerController {
	return &ServerController{
		UserController: userController,
	}
}