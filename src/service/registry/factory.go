package registry

import (
	"members/common"
	"members/config"
	"members/service"
	"members/storage"

	"go.uber.org/fx"
)

type (
	registryFactory struct{}
	RegistryFactory = service.ServiceFactory[*registryService]
)

var (
	Module = fx.Module(
		"registry-service",
		fx.Provide(
			fx.Annotate(
				create,
				fx.As(new(RegistryFactory)),
			),
		),
		fx.Invoke(
			service.Create[*registryService](common.ServiceRegistry),
		),
	)
)

var (
	_ RegistryFactory = ((*registryFactory)(nil))
)

func create(svc *service.SvcFramework) *registryFactory {
	return &registryFactory{}
}

func (h *registryFactory) CreateService(cfg config.ConfigProvider, store storage.Store) *registryService {
	return &registryService{
		store: store,
	}
}
