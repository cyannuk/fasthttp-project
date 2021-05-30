//+build wireinject

package composition

import (
	"fasthttp-project/api"
	"fasthttp-project/config"
	"fasthttp-project/domain/repository"
	"fasthttp-project/domain/service"
	"github.com/google/wire"
)

func Application() (*api.Application, error) {
	wire.Build(config.GetDatabaseConfig, config.GetServerConfig, repository.NewDataSource, repository.NewUserRepository, repository.NewOrderRepository,
		service.NewUserService, service.NewOrderService, api.NewApplication)
	return &api.Application{}, nil
}
