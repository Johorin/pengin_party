package usecase

import (
	"context"
	"pengin_party/internal/domain/user"
	"pengin_party/internal/infrastructure/repository"
)

type IsExistUserUseCase struct {
	db       repository.DBInterface
	userRepo user.UserRepository
}

func NewIsExistUserUseCase(
	db repository.DBInterface,
	userRepo user.UserRepository,
) *IsExistUserUseCase {
	return &IsExistUserUseCase{
		db,
		userRepo,
	}
}

func (uc *IsExistUserUseCase) Execute(
	ctx context.Context,
	uid string,
) (bool, error) {
	isExist, err := uc.userRepo.Search(ctx, uid)
	if err != nil {
		return false, err
	}
	return isExist, err
}