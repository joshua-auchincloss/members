package admin

import (
	"members/common"
	"members/config"
	"members/service"
	"members/service/core"
	"members/storage/base"
)

type (
	AdminFactory = service.ServiceFactory[*adminService]
)

var (
	Module = core.NewModule[
		AdminFactory, *adminService,
	](
		"admin-service",
		common.ServiceAdmin,
		New,
	)
)

func New() AdminFactory {
	return core.New[*adminService](
		func(cfg config.ConfigProvider, store base.BaseStore) *adminService {
			return &adminService{
				store: store,
			}
		},
	)
}
