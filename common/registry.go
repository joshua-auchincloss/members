package common

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type (
	RegisteredProto struct {
		bun.BaseModel `bun:"registered_protobufs"`

		Id         string    `bun:"id"`
		FileName   string    `bun:"file"`
		Data       []byte    `bun:"data"`
		LastUpdate time.Time `bun:"last_update,default:current_timestamp"`
	}

	ProtoMeta struct {
		bun.BaseModel `bun:"registration_meta"`

		Id           string    `bun:"id,pk"`
		Key          string    `bun:"key"`
		Version      string    `bun:"version"`
		RegisteredAt time.Time `bun:"registration,default:current_timestamp"`
	}

	ProtoProject struct {
		bun.BaseModel `bun:"repos"`

		Id       string    `bun:"id,pk"`
		Name     string    `bun:"name"`
		Owner    string    `bun:"owner"`
		Creation time.Time `bun:"creation,default:current_timestamp"`

		Meta []*ProtoMeta `bun:"rel:has-many,join:id=key"`
	}
)

func (*ProtoMeta) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		Model((*ProtoMeta)(nil)).
		Index("proto_meta_keyed_idx").
		Column("key").
		Exec(ctx)
	return err
}
