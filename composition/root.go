package composition

import (
	"fasthttp-project/api"
	"fasthttp-project/config"
	"fasthttp-project/domain/repository"
	"fasthttp-project/domain/service"
	"go.uber.org/dig"
)

func compose() (*dig.Container, error) {
	var container = dig.New()
	err := container.Provide(config.GetDatabaseConfig)
	if err != nil {
		return nil, err
	}
	err = container.Provide(config.GetServerConfig)
	if err != nil {
		return nil, err
	}
	err = container.Provide(repository.NewDataSource)
	if err != nil {
		return nil, err
	}
	err = container.Provide(repository.NewUserRepository)
	if err != nil {
		return nil, err
	}
	err = container.Provide(repository.NewOrderRepository)
	if err != nil {
		return nil, err
	}
	err = container.Provide(service.NewUserService)
	if err != nil {
		return nil, err
	}
	err = container.Provide(service.NewOrderService)
	if err != nil {
		return nil, err
	}
	err = container.Provide(api.NewApplication)
	if err != nil {
		return nil, err
	}
	return container, nil
}

func Application() error {
	container, err := compose()
	if err != nil {
		return err
	}
	return container.Invoke(func(application *api.Application, config *config.ServerConfig) error {
		return application.Start(config.BindAddress())
	})
}
