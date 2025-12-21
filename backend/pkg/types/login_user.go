package types

type LoginUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	UID   string `json:"uid"`
	UUID  string `json:"uuid"`
}