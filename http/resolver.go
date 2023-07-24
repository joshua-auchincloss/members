package server

import (
	"members/common"
	"members/config"
	"members/storage"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/resolver"
)

type (
	resolverBuilder struct {
		kind  common.Service
		store storage.Store
		prov  config.ConfigProvider
	}

	grpcresolver struct {
		target resolver.Target
		cc     resolver.ClientConn

		kind  common.Service
		prov  config.ConfigProvider
		store storage.Store
	}
)

func (b *resolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &grpcresolver{
		target: target,
		cc:     cc,
		kind:   b.kind,
		store:  b.store,
		prov:   b.prov,
	}
	r.start()
	return r, nil
}
func (b *resolverBuilder) Scheme() string { return common.ServiceKeys.Get(b.kind) }

func (r *grpcresolver) start() {
	dyn := r.prov.GetDynamic()
	upd := func() {
		toupd := []resolver.Address{}
		for _, addr := range dyn.GetDns(r.kind, r.target.Endpoint()) {
			toupd = append(toupd, resolver.Address{Addr: addr})
		}
		log.Info().Str("endpoint", r.target.Endpoint()).Interface("target", r.target).Interface("addrs", toupd).Send()
		r.cc.UpdateState(resolver.State{Addresses: toupd})
	}
	go func() {
		for range dyn.Subscription() {
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
