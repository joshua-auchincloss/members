package server

import (
	"members/common"
	"members/config"
	storage "members/storage/base"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc/resolver"
)

func LoadBalancerFor(svc common.Service) fx.Option {
	return fx.Invoke(func(ctx *cli.Context, prov config.ConfigProvider, store storage.BaseStore) error {
		resolver.Register(&resolverBuilder{svc, store, prov})
		return nil
	})
}
