package p2p

import (
	"context"
	"members/common"
	"members/config"
	"members/storage"
)

type (
	P2PRegistry struct {
		prov  config.ConfigProvider
		store storage.Store
	}

	participants = []string
)

func (h *P2PRegistry) GetStore() storage.Store {
	return h.store
}

func newRegistry(prov config.ConfigProvider, store storage.Store) *P2PRegistry {
	return &P2PRegistry{
		prov,
		store,
	}
}
func (h *P2PRegistry) ProposeHandler(ctx context.Context, meta *common.ProtoMeta, proto *common.RegisteredProto) error {
	return h.store.RegisterProto(ctx, meta, proto)
}
