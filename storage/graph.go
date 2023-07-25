package storage

import (
	"context"
	"members/common"
	"members/config"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type (
	Graph struct {
		db     *dgo.Dgraph
		logger *zerolog.Logger
	}
	GraphInitializer = func(prov config.ConfigProvider) (*dgo.Dgraph, error)
)

func NewGraph(db *dgo.Dgraph) Store {
	return &Graph{db, nil}
}

var (
	_ Store = ((*Graph)(nil))
)

func (gr *Graph) CleanOldMembers(ctx context.Context, from time.Duration) error {
	// const q = `
	// {
	// 	all(func: anyofterms(name, "Alice Bob")) {
	// 		uid
	// 		balance
	// 	}
	// }
	// `
	// resp, err := sq.db.NewTxn().Query(context.Background(), q)
	// if err != nil {
	// 	panic(err)
	// 	// return err
	// }
	// sq.logger.Info().Interface("re", resp).Send()
	return nil
}

func (gr *Graph) WithLogger(sub *zerolog.Logger) {
	gr.logger = sub
}

func (gr *Graph) UpsertMembership(ctx context.Context, meta *common.Membership) error {
	return nil
}

func (gr *Graph) GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error) {
	var out []*common.Membership

	return out, nil
}

func (gr *Graph) CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error {
	project.Id = uuid.NewString()
	proto.Id = uuid.NewString()
	proto.Key = project.Id
	return nil
}

func (gr *Graph) Teardown() error {
	return nil
}

func (gr *Graph) create() error {
	return nil
}

func (gr *Graph) Setup(cfg config.ConfigProvider) error {
	storg := cfg.GetConfig().Storage
	if storg.Drop {
	}
	if storg.Create {
	}
	op := &api.Operation{}
	return gr.db.Alter(context.TODO(), op)
}

func (gr *Graph) Registered(key string) bool {
	return false
}

func (gr *Graph) GetHandler(key string) (*common.RegisteredProto, error) {
	return &common.RegisteredProto{}, nil
}

func (gr *Graph) RegisterProto(ctx context.Context, proto *common.ProtoMeta, data *common.RegisteredProto) error {
	id := uuid.NewString()
	proto.RegisteredAt = time.Now().UTC()
	proto.Id = id
	data.Id = id
	return nil

}
