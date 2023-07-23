package service

import (
	"members/common"
	"members/config"
	"members/utils"

	"github.com/rs/zerolog/log"
)

var (
	service_keys = map[common.Service]string{
		common.ServiceHealth:   "health",
		common.ServiceRegistry: "registry",
	}
)

func ForService(prov config.ConfigProvider, key common.Service) bool {
	svcs := prov.GetConfig().Services
	log.Print(svcs)
	if len(svcs) != 0 {
		if svcs[0] == "all" {
			return true
		}
		return utils.AnyEq(svcs, service_keys[key])
	}
	return utils.AnyEq(svcs, service_keys[key])
}
