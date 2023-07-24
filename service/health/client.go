package health

import (
	"members/common"
	"members/grpc/api/v1/health"
	"members/service"

	"go.uber.org/fx"
)

type (
	HealthClient = service.ClientFactory[health.HealthClient]
)

var (
	client_factory = service.NewClientFactory(common.ServiceHealth, health.NewHealthClient)
	ClientFactory  = fx.Module("health-client-factory",
		fx.Supply(client_factory),
	)
)
