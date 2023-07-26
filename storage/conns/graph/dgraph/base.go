package dgraph

import (
	"context"
	"members/common"
	"members/config"
	"members/storage/base"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/rs/zerolog"
)

type (
	DgraphBase struct {
		DB     *dgo.Dgraph
		logger *zerolog.Logger
		kind   common.StorageType
		prov   config.ConfigProvider
	}
	GraphInitializer = func(prov config.ConfigProvider) (*dgo.Dgraph, error)
	GraphConverter   = func(*dgo.Dgraph) (base.BaseStore, error)
)

var (
	_ base.BaseStore = ((*DgraphBase)(nil))
)

func Invoker(
	prov config.ConfigProvider,
	factory base.StoreFactory,
	root *zerolog.Logger,
) (base.BaseStore, error) {
	sub := base.LoggerFor(prov, root)
	if gr, err := factory.New(); err != nil {
		return nil, err
	} else {
		gr.WithLogger(sub)
		gr.WithProvider(prov)
		return gr, nil
	}
}

func NewDgraphBase(db *dgo.Dgraph, log *zerolog.Logger, typ common.StorageType) (base.BaseStore, error) {
	return &DgraphBase{db, log, common.StorageDGraph, nil}, nil
}

func (gr *DgraphBase) Kind() common.StorageType {
	return gr.kind
}

func (gr *DgraphBase) WithQueryHook(i interface{}) {
}

func (gr *DgraphBase) WithProvider(prov config.ConfigProvider) {
}

func (gr *DgraphBase) WithLogger(sub *zerolog.Logger) {
	gr.logger = sub
}

func (gr *DgraphBase) Logger() *zerolog.Logger {
	return gr.logger
}

func (gr *DgraphBase) UpsertMembership(ctx context.Context, meta *common.Membership) error {
	return nil
}

func (gr *DgraphBase) CleanOldMembers(ctx context.Context, from time.Duration) error {
	return nil
}

func (gr *DgraphBase) GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error) {
	var out []*common.Membership

	return out, nil
}

func (gr *DgraphBase) Teardown(ctx context.Context) error {
	return nil
}

func (gr *DgraphBase) Setup(ctx context.Context, cfg config.ConfigProvider) error {
	storg := cfg.GetConfig().Storage
	if storg.Drop {
	}
	if storg.Create {
	}
	op := &api.Operation{}
	return gr.DB.Alter(context.TODO(), op)
}
