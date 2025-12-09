package usecase

type UserInfo struct {
	ID   uint
	UID  uint
	UUID uint
	Name string
}

type CreateUserResponseData struct {
	ID   uint   `json:"id"`
	UID  uint   `json:"uid"`
	UUID uint   `json:"uuid"`
	Name string `json:"Name"`
}

type CreateUserResponse struct {
	Data CreateUserResponseData `json:"data"`
}