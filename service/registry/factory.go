package registry

import (
	"members/common"
	"members/config"
	"members/service"
	"members/service/core"
	"members/storage/base"
)

type (
	RegistryFactory = service.ServiceFactory[*registryService]
)

var (
	Module = core.NewModule[
		RegistryFactory, *registryService,
	](
		"registry",
		common.ServiceRegistry,
		New,
	)
)

func New() RegistryFactory {
	return core.New[*registryService](
		func(cfg config.ConfigProvider, store base.BaseStore) *registryService {
			return &registryService{
				// store: store,
			}
		},
	)
}

// func (h *registryFactory) CreateService(cfg config.ConfigProvider, store storage.Store) *registryService {
// 	return &registryService{
// 		store: store,
// 	}
// }
