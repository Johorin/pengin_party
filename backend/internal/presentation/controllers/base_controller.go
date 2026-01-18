package controllers

type ServerController struct {
	UserController UserController
	RoomController RoomController
}

func NewServerController(
	userController UserController,
	roomController RoomController,
) *ServerController {
	return &ServerController{
		UserController: userController,
		RoomController: roomController,
	}
}