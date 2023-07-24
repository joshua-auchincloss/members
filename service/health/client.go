package health

import (
	"members/common"
	"members/grpc/api/v1/health"
	server "members/http"
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
		server.LoadBalancerFor(common.ServiceHealth),
	)
)
