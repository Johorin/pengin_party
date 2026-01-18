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
	Name string `json:"name"`
}

type UserApiResponse struct {
	Data UserResponse `json:"data"`
}

type CreateUserResponse struct {
	ID *uint `json:"id"`
}

type CreateRoomResponse struct {
	RoomID *string `json:"room_id"`
}

type SearchUserResponse struct {
	IsExist bool `json:"is_exist"`
}

type CreateUserApiResponse struct {
	Data CreateUserResponse `json:"data"`
}

type CreateRoomApiResponse struct {
	Data CreateRoomResponse `json:"data"`
}

type SearchUserApiResponse struct {
	Data SearchUserResponse `json:"data"`
}