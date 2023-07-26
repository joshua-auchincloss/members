package base

import (
	"context"
	"members/common"
	"members/config"
	"time"

	"github.com/rs/zerolog"
)

const ()

type (
	StorageType = int

	BaseStore interface {
		Kind() common.StorageType
		WithLogger(sub *zerolog.Logger)
		WithProvider(prov config.ConfigProvider)

		Setup(ctx context.Context, config config.ConfigProvider) error
		Teardown(ctx context.Context) error

		Logger() *zerolog.Logger
		WithQueryHook(interface{})

		GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error)
		CleanOldMembers(ctx context.Context, from time.Duration) error
		UpsertMembership(ctx context.Context, meta *common.Membership) error
	}

	StoreFactory interface {
		New() (BaseStore, error)
	}
)
