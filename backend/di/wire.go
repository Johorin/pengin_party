package di

import (
	userUC "pengin_party/internal/application/usecases/user/usecase"
	controller "pengin_party/internal/presentation/controllers"

	"github.com/google/wire"
)

var useCaseSet = wire.NewSet(
	userUC.NewCreateUserUseCase,
)

var serverControllerSet = wire.NewSet(
	controller.NewUserController,
	controller.NewServerController,
)

type ControllerSet struct {
	ServerController *controller.ServerController
	// DB               repository.DBInterface
	// Cache            cache.CacheRepository
}

func InitializeControllers() (*ControllerSet, error) {
	wire.Build(
		useCaseSet,
		serverControllerSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil, nil
}
