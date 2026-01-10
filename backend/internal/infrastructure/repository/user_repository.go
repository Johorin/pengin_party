package repository

import (
	"context"
	"fmt"
	"pengin_party/internal/domain/user"
	"pengin_party/internal/infrastructure/dbmodel"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db DBInterface
}

func NewUserRepository(db DBInterface) user.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUserIfNotExists(ctx context.Context, entity *user.UserEntity) (*uint, error) {
	model := r.toModel(entity)
	result := r.db.GetDB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}}, // 衝突判定フィールド
		DoNothing: true,                             // 既存なら INSERT しない
	}).Create(&model)

	if result.Error != nil {
		return nil, fmt.Errorf("userの作成に失敗しました: %w", result.Error)
	}
	// 既に同じemailのユーザーが存在していたらカスタムエラーを返す
	if result.RowsAffected == 0 {
		return nil, &user.ErrUserAlreadyExists{Email: model.Email}
	}

	return &model.ID, nil
}

func (r *userRepository) Search(ctx context.Context, uid string) (bool, error) {
	var isExist bool
	var model dbmodel.User
	result := r.db.GetDB().WithContext(ctx).Where("uid = ?", uid).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			isExist = false
		}
	} else {
		isExist = true
	}
	return isExist, nil
}

func (r *userRepository) toModel(entity *user.UserEntity) *dbmodel.User {
	return &dbmodel.User{
		Name:  entity.Name(),
		Email: entity.Email().Value(),
		UID:   entity.UID(),
		UUID:  entity.UUID(),
	}
}
