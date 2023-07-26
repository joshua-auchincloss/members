package health

import (
	"members/common"
	"members/config"
	"members/service"
	"members/service/core"
	"members/storage/base"
	"time"
)

type (
	ServiceFactory = service.ServiceFactory[*healthService]
)

var (
	Module = core.NewModule[
		ServiceFactory, *healthService,
	](
		"health-service",
		common.ServiceHealth,
		New,
	)
)

var (
	default_polling = time.Second * 5
)

func New() ServiceFactory {
	return core.New[*healthService](
		func(cfg config.ConfigProvider, store base.BaseStore) *healthService {
			return &healthService{
				store: store,
			}
		},
	)
}
