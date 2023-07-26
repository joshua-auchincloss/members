package admin

import (
	"members/common"
	"members/config"
	"members/service"
	"members/storage/base"

	"go.uber.org/fx"
)

type (
	adminFactory struct{}
	AdminFactory = service.ServiceFactory[*adminService]
)

var (
	// Module = service.GrpcModule[adminService, registry.UnimplementedRegistryHandler](common.ServiceRegistry, registryconnect.NewRegistryHandler)

	Module = fx.Module(
		"registry-service",
		fx.Provide(
			fx.Annotate(
				create,
				fx.As(new(AdminFactory)),
			),
		),
		fx.Invoke(
			service.Create[*adminService](common.ServiceAdmin),
		),
	)
)

var (
	_ AdminFactory = ((*adminFactory)(nil))
)

func create(svc *service.SvcFramework) *adminFactory {
	return &adminFactory{}
}

func (h *adminFactory) CreateService(cfg config.ConfigProvider, store base.BaseStore) *adminService {
	return &adminService{
		store: store,
	}
}
