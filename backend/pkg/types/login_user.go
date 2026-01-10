package types

type LoginUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	UID   string `json:"uid"`
}

type UserIDResponse struct {
	UserID *uint `json:"user_id"`
}