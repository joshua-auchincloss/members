package admin

import (
	"members/common"
	"members/grpc/api/v1/admin"
	"members/server"
	"members/service"

	"go.uber.org/fx"
)

type (
	AdminClient = *service.ClientFactory[admin.AdminClient]
)

var (
	client_factory = service.NewClientFactory(common.ServiceAdmin, admin.NewAdminClient)

	ClientFactory = fx.Module("admin-client-factory",
		fx.Supply(client_factory),
		server.LoadBalancerFor(common.ServiceAdmin),
	)
)
