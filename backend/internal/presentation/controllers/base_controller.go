package controllers

type ServerController struct {
	UserController UserController
	RoomController RoomController
	WebSocketController WebSocketController
}

func NewServerController(
	userController UserController,
	roomController RoomController,
	webSocketController WebSocketController,
) *ServerController {
	return &ServerController{
		UserController: userController,
		RoomController: roomController,
		WebSocketController: webSocketController,
	}
}