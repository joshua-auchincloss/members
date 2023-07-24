package server

import (
	"members/config"
	"members/storage"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/resolver"
)

type (
	resolverBuilder struct {
		store storage.Store
		prov  config.ConfigProvider
	}

	grpcresolver struct {
		target resolver.Target
		cc     resolver.ClientConn

		prov  config.ConfigProvider
		store storage.Store
	}
)

func (b *resolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &grpcresolver{
		target: target,
		cc:     cc,
		store:  b.store,
		prov:   b.prov,
	}
	r.start()
	return r, nil
}
func (*resolverBuilder) Scheme() string { return lb_scheme }

func (r *grpcresolver) start() {
	dyn := r.prov.GetDynamic()
	upd := func() {
		toupd := []resolver.Address{}
		for _, addr := range dyn.GetDns(r.target.Endpoint()) {
			toupd = append(toupd, resolver.Address{Addr: addr})
		}
		log.Info().Str("endpoint", r.target.Endpoint()).Interface("target", r.target).Interface("addrs", toupd).Send()
		r.cc.UpdateState(resolver.State{Addresses: toupd})
	}
	go func() {
		for {
			time.Sleep(time.Second * 5)
			upd()
		}
	}()
	upd()
}
func (*grpcresolver) ResolveNow(o resolver.ResolveNowOptions) {
	log.Info().Interface("resolver", o).Send()
}
func (*grpcresolver) Close() {
	log.Info().Str("resolver", "close").Send()
}
