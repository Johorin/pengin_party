package usecase

import "context"

type CreateUserUseCase struct {}

func NewCreateUserUseCase() *CreateUserUseCase {
	return &CreateUserUseCase{}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, user UserInfo) (*CreateUserResponse, error) {
	// TODO: repository処理、usersテーブルに対してInsert

	return &CreateUserResponse{
		Data: CreateUserResponseData{
			ID: 1,
			UID: 11,
			UUID: 111,
			Name: "yuya",
		},
	}, nil
}