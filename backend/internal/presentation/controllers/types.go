package controllers

type RoomResponse struct {
	RoomId string `json:"roomId"`
	Status string `json:"status"`
}

type RoomApiResponse struct {
	Data RoomResponse `json:"data"`
}

type UserResponse struct {
	ID uint `json:"id"`
	UID uint `json:"uid"`
	UUID uint `json:"uuid"`
	Name string `json:"name"`
}

type UserApiResponse struct {
	Data UserResponse `json:"data"`
}

type CreateUserResponse struct {
	ID *uint `json:"id"`
}

type CreateUserApiResponse struct {
	Data CreateUserResponse `json:"data"`
}