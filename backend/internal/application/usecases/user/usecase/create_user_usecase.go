package usecase

import (
	"context"
	"pengin_party/internal/domain/user"
	"pengin_party/internal/infrastructure/repository"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

const CreateUserUseCaseMessage = "ユーザーの作成に成功しました"

type CreateUserUseCase struct {
	db       repository.DBInterface
	userRepo user.UserRepository
}

func NewCreateUserUseCase(
	db repository.DBInterface,
	userRepo user.UserRepository,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		db:       db,
		userRepo: userRepo,
	}
}

func (uc *CreateUserUseCase) Execute(
	ctx context.Context,
	name, email, uid, uuid string,
) (*uint, error) {
	// Userエンティティの作成
	userEntity, err := user.NewUserEntity(
		name,
		email,
		uid,
		uuid,
	)
	if err != nil {
		return nil, errors.Wrap(err, "ユーザーエンティティの作成に失敗しました")
	}

	var userId *uint
	if err := uc.db.Transaction(ctx, func(tx *gorm.DB) error {
		var err error
		userId, err = uc.userRepo.CreateUserIfNotExists(ctx, userEntity)
		if err != nil {
			return errors.Wrap(err, "ユーザーの作成に失敗しました")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return userId, nil
}