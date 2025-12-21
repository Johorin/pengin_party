package repository

import (
	"context"
	"fmt"
	"pengin_party/internal/domain/user"
	"pengin_party/internal/infrastructure/dbmodel"
)

type userRepository struct {
	db DBInterface
}

func NewUserRepository(db DBInterface) user.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, entity *user.UserEntity) (*uint, error) {
	model := r.toModel(entity)
	if err := r.db.GetDB().WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("userの作成に失敗しました: %w", err)
	}
	return &model.ID, nil
}

func (r *userRepository) toModel(entity *user.UserEntity) *dbmodel.User {
	return  &dbmodel.User{
		Name:  entity.Name(),
		Email: entity.Email().Value(),
		UID:   entity.UID(),
		UUID:  entity.UUID(),
	}
}