package p2p

import (
	"members/common"
	"members/config"
	"members/storage"

	"github.com/uptrace/bun"
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
func (h *P2PRegistry) ProposeHandler(meta *common.ProtoMeta, proto *common.RegisteredProto) error {
	return h.store.RegisterProto(meta, proto)
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
}
