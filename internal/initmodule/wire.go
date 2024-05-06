//go:build wireinject
// +build wireinject

package initmodule

import (
	"context"

	"coincap/internal/facade"
	"coincap/internal/gateway"
	"coincap/internal/repository"
	"coincap/internal/rest_api/controller"
	"coincap/internal/usecase"
	"coincap/pkg/cfg"
	"coincap/pkg/database"

	"github.com/google/wire"
)

var (
	pkgSet = wire.NewSet(
		database.Init,
	)

	gatewaySet = wire.NewSet(
		wire.Struct(new(gateway.CoinCapGatewayParams), "*"),
		gateway.NewCoinCapGateway,
	)

	repositorySet = wire.NewSet(
		wire.Struct(new(repository.UserRepositoryParams), "*"),
		repository.NewUserRepository,

		wire.Struct(new(repository.MyCoinRepositoryParams), "*"),
		repository.NewMyCoinRepository,
	)

	usecaseSet = wire.NewSet(
		wire.Struct(new(usecase.UserUseCaseParams), "*"),
		usecase.NewUserUseCase,

		wire.Struct(new(usecase.CoinCapUseCaseParams), "*"),
		usecase.NewCoinUsecase,

		wire.Struct(new(usecase.MyCoinUseCaseParams), "*"),
		usecase.NewMyCoinUsecase,

		usecase.New,
	)

	facadeSet = wire.NewSet(
		wire.Struct(new(facade.UserFacadeParams), "*"),
		facade.NewUserFacade,

		wire.Struct(new(facade.CoinFacadeParams), "*"),
		facade.NewCoinFacade,

		wire.Struct(new(facade.MyCoinFacadeParams), "*"),
		facade.NewMyCoinFacade,
	)

	controllerSet = wire.NewSet(
		controller.NewHealthCheck,

		wire.Struct(new(controller.MyCoinCtrlParams), "*"),
		controller.NewMyCoinCtrl,

		wire.Struct(new(controller.UserCtrlParams), "*"),
		controller.NewUserCtrl,

		wire.Struct(new(controller.CoinCtrlParams), "*"),
		controller.NewCoinCtrl,
		controller.New,
	)

	restApiSet = wire.NewSet(
		NewRestAPIWrapper,
	)

	restSet = wire.NewSet(
		pkgSet,
		gatewaySet,
		repositorySet,
		usecaseSet,
		facadeSet,
		controllerSet,
		restApiSet,
	)
)

func InitAppServer(ctx context.Context, config *cfg.ConfigSchema) *RestAPIWrapper {
	wire.Build(restSet)
	return new(RestAPIWrapper)
}
