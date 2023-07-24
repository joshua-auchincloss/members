package service

import (
	"members/common"
	"members/config"
	"members/utils"

	"github.com/rs/zerolog/log"
)

func ForService(prov config.ConfigProvider, key common.Service) bool {
	svcs := prov.GetConfig().Services
	log.Print("SERVICES", svcs)
	if len(svcs) != 0 {
		if svcs[0] == "all" {
			return true
		}
		return utils.AnyEq(svcs, common.ServiceKeys.Get(key))
	}
	return false
}
