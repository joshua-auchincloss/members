package storage

import (
	"members/common"
	"members/config"
)

type (
	Store interface {
		Setup(config config.ConfigProvider) error
		Registered(key string) bool
		GetHandler(key string) (*common.RegisteredProto, error)
		UpsertMembership(meta *common.Membership) error
		GetMembers(kind ...common.Service) ([]*common.Membership, error)
		RegisterProto(proto *common.ProtoMeta, data *common.RegisteredProto) error
	}

	StorageType = int
)

func Setup(config config.ConfigProvider, store Store) error {
	return store.Setup(config)
}

const (
	Memory StorageType = iota
	Postgres
)
