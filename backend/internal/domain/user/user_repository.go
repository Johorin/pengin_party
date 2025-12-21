package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, entity *UserEntity) (*uint, error)
}