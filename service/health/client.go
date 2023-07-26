package health

import (
	"members/common"
	"members/grpc/api/v1/health"
	"members/service"
	"members/service/core"
)

type (
	HealthClient = service.ClientFactory[health.HealthClient]
)

var (
	ClientFactory = core.NewClient(
		"health",
		common.ServiceHealth,
		health.NewHealthClient,
	)
)
