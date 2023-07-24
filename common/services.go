package common

import (
	commonpb "members/grpc/api/v1/common"
)

type (
	Service = commonpb.Service
)

const (
	UnknownService  Service = commonpb.Service_SERVICE_UNKNOWN
	ServiceHealth           = commonpb.Service_SERVICE_HEALTH
	ServiceRegistry         = commonpb.Service_SERVICE_REGISTRY
	ServiceAdmin            = commonpb.Service_SERVICE_ADMIN
)

var (
	ServiceKeys = NewKey(
		map[Service]string{
			ServiceHealth:   "health",
			ServiceRegistry: "registry",
			ServiceAdmin:    "admin",
		},
	)
	ServiceNames = ServiceKeys.Reverse()
)
