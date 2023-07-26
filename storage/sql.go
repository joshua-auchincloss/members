package storage

import (
	"context"
	"database/sql"
	"members/common"
	sqlbase "members/storage/conns/sql"
	"time"

	"github.com/google/uuid"
)

type (
	Sql struct {
		sqlbase.SqlBase
	}
)

var (
	_ Store = ((*Sql)(nil))
)

func NewSql(base sqlbase.SqlBase) Store {
	return &Sql{
		base,
	}
}

func (sq *Sql) CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error {
	project.Id = uuid.NewString()
	proto.Id = uuid.NewString()
	proto.Key = project.Id
	if _, err := sq.DB.NewInsert().
		Model(project).
		Exec(ctx, project); err != nil && err != sql.ErrNoRows {
		sq.Logger().Print(err)
		return err
	}
	if _, err := sq.DB.NewInsert().
		Model(proto).
		Exec(ctx, proto); err != nil && err != sql.ErrNoRows {
		sq.Logger().Print(err)
		return err
	}
	return nil
}

func (sq *Sql) Registered(key string) bool {
	return false
}

func (sq *Sql) GetHandler(key string) (*common.RegisteredProto, error) {
	return &common.RegisteredProto{}, nil
}

func (sq *Sql) RegisterProto(ctx context.Context, proto *common.ProtoMeta, data *common.RegisteredProto) error {
	id := uuid.NewString()
	proto.RegisteredAt = time.Now().UTC()
	proto.Id = id
	data.Id = id
	return nil

}
