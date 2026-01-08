package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, entity *UserEntity) (*uint, error)
	Search(ctx context.Context, uid string) (bool, error)
}