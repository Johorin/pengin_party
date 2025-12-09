package controllers

type ServerController struct {
	userController UserController
}

func NewServerController(
	userController UserController,
) *ServerController {
	return &ServerController{
		userController: userController,
	}
}