//go:build wireinject
// +build wireinject

package di

import (
	userUC "pengin_party/internal/application/usecases/user/usecase"
	"pengin_party/internal/presentation/controllers"
	"pengin_party/internal/infrastructure/repository"
	"pengin_party/internal/infrastructure/redis"

	"github.com/google/wire"
)

var infrastructureSet = wire.NewSet(
	repository.Init,
	redis.NewRedisClient,
)

var repositorySet = wire.NewSet(
	repository.NewUserRepository,
)

var useCaseSet = wire.NewSet(
	userUC.NewCreateUserUseCase,
	userUC.NewIsExistUserUseCase,
)

var controllerSet = wire.NewSet(
	controllers.NewUserController,
)

var serverControllerSet = wire.NewSet(
	controllers.NewServerController,
)

type ControllerSet struct {
	ServerController *controllers.ServerController
	DB               repository.DBInterface
	// Cache            cache.CacheRepository
}

func InitializeControllers() (*ControllerSet, error) {
	wire.Build(
		infrastructureSet,
		repositorySet,
		useCaseSet,
		controllerSet,
		serverControllerSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil, nil
}
