package storage

import (
	"context"
	"database/sql"
	"log"
	"members/common"
	"members/config"
	"members/utils"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Sql struct {
		db *bun.DB
	}
	Initializer = func(prov config.ConfigProvider) (*bun.DB, error)
)

var (
	_ Store = ((*Sql)(nil))
)

func NewSql(db *bun.DB) Store {
	return &Sql{db}
}

func runCreate[T interface{}](
	sq *Sql,
	model T,
	chain ...func(*bun.CreateTableQuery) *bun.CreateTableQuery,
) error {
	var fc func(*bun.CreateTableQuery) *bun.CreateTableQuery
	if len(chain) != 0 {
		fc = chain[0]
	}
	base := sq.db.NewCreateTable().
		Model(model).
		IfNotExists()
	if fc != nil {
		base = fc(base)
	}
	if _, err := base.Exec(context.TODO()); err != nil {
		return err
	}
	return nil
}

func drop[T interface{}](sq *Sql, model T) {
	sq.db.NewDropTable().
		Model(model).Exec(context.TODO())
}

func (sq *Sql) UpsertMembership(meta *common.Membership) error {
	if _, err := sq.db.NewInsert().
		Model(meta).
		On(
			"conflict (address, service) do update set last_health = excluded.last_health",
		).Exec(context.TODO(), meta); err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (sq *Sql) GetMembers(kind ...common.Service) ([]*common.Membership, error) {
	ok := utils.OkSpread(kind)
	var out []*common.Membership
	base := sq.db.NewSelect().
		Model(&out)
	if ok {
		base = base.Where(
			"service in (?)", bun.In(kind),
		)
	}
	if err := base.Scan(context.TODO(), &out); err != nil {
		log.Print(err)
		return nil, err
	}
	return out, nil
}

func (sq *Sql) Setup(cfg config.ConfigProvider) error {
	if cfg.GetConfig().Storage.Drop {
		for _, tbl := range []interface{}{
			(*common.Membership)(nil),
			(*common.ProtoProject)(nil),
			(*common.ProtoMeta)(nil),
			(*common.RegisteredProto)(nil),
		} {
			drop(sq, tbl)
		}
	}

	if err := runCreate(sq, (*common.Membership)(nil)); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.ProtoProject)(nil)); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.ProtoMeta)(nil), func(ctq *bun.CreateTableQuery) *bun.CreateTableQuery {
		return ctq.ForeignKey(`("project_key") REFERENCES "repos" ("id") ON DELETE CASCADE`)
	}); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.RegisteredProto)(nil), func(ctq *bun.CreateTableQuery) *bun.CreateTableQuery {
		return ctq.ForeignKey(`("id") REFERENCES "registration_meta" ("id") ON DELETE CASCADE`)
	}); err != nil {
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

func (sq *Sql) RegisterProto(proto *common.ProtoMeta, data *common.RegisteredProto) error {
	id := uuid.NewString()
	proto.RegisteredAt = time.Now().UTC()
	proto.Id = id
	data.Id = id
	return nil

}
