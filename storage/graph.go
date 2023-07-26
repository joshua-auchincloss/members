package storage

import (
	"context"
	"members/common"
	"members/config"
	"members/storage/conns/graph/dgraph"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/google/uuid"
)

type (
	Graph struct {
		*dgraph.DgraphBase
	}
	GraphInitializer = func(prov config.ConfigProvider) (*dgo.Dgraph, error)
)

var (
	_ Store = ((*Graph)(nil))
)

func (gr *Graph) CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error {
	project.Id = uuid.NewString()
	proto.Id = uuid.NewString()
	proto.Key = project.Id
	return nil
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
