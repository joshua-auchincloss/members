package starter

import (
	"members/config"
	"members/service"

	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"service-starter",
		fx.Invoke(start),
	)
)

func start(lc fx.Lifecycle, svc *service.SvcFramework, prov config.ConfigProvider) {
	svc.Start(lc, prov)
}
