package usecase

import (
	"context"
	"pengin_party/internal/domain/user"
	"pengin_party/internal/infrastructure/repositories/rdb"
)

type IsExistUserUseCase struct {
	db       rdb.DBInterface
	userRepo user.UserRepository
}

func NewIsExistUserUseCase(
	db rdb.DBInterface,
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