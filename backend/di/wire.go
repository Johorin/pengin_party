//go:build wireinject
// +build wireinject

package di

import (
	roomUC "pengin_party/internal/application/usecases/room/usecase"
	userUC "pengin_party/internal/application/usecases/user/usecase"
	"pengin_party/internal/infrastructure/repositories/rdb"
	rdbRepo "pengin_party/internal/infrastructure/repositories/rdb/repository"
	"pengin_party/internal/infrastructure/repositories/redis"
	redisRepo "pengin_party/internal/infrastructure/repositories/redis/repository"
	"pengin_party/internal/presentation/controllers"
	"pengin_party/internal/infrastructure/firebase"

	"github.com/google/wire"
)

var infrastructureSet = wire.NewSet(
	rdb.Init,
	redis.Init,
	firebase.Init,
)

var rdbRepositorySet = wire.NewSet(
	rdbRepo.NewUserRepository,
)

var redisRepositorySet = wire.NewSet(
	redisRepo.NewRoomRepository,
)

var useCaseSet = wire.NewSet(
	userUC.NewCreateUserUseCase,
	userUC.NewIsExistUserUseCase,
	roomUC.NewCreateRoomUseCase,
	roomUC.NewJoinRoomUseCase,
)

var controllerSet = wire.NewSet(
	controllers.NewUserController,
	controllers.NewRoomController,
)

var serverControllerSet = wire.NewSet(
	controllers.NewServerController,
)

type ControllerSet struct {
	ServerController *controllers.ServerController
	DB               rdb.DBInterface
	Redis            redis.RedisInterface
	Firebase         firebase.FirebaseInterface
	// Cache            cache.CacheRepository
}

func InitializeControllers() (*ControllerSet, error) {
	wire.Build(
		infrastructureSet,
		rdbRepositorySet,
		redisRepositorySet,
		useCaseSet,
		controllerSet,
		serverControllerSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil, nil
}
