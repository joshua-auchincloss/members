package admin

import (
	"members/common"
	"members/grpc/api/v1/admin"
	"members/service"
	"members/service/core"
)

type (
	AdminClient = *service.ClientFactory[admin.AdminClient]
)

var (
	ClientFactory = core.NewClient(
		"admin",
		common.ServiceAdmin,
		admin.NewAdminClient,
	)
)
