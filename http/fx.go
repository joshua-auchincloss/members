package server

import (
	"members/config"
	"members/storage"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc/resolver"
)

var (
	Lb = fx.Module(
		"server-loadbalance",
		fx.Invoke(ensure_lb),
	)
)

func ensure_lb(ctx *cli.Context, prov config.ConfigProvider, store storage.Store) error {
	// addrs := ctx.StringSlice(config.RemoteAddrsDefn.Key)
	resolver.Register(&resolverBuilder{store, prov})
	return nil
}
