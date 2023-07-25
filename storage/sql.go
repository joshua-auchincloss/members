package storage

import (
	"context"
	"database/sql"
	"fmt"
	"members/common"
	"members/config"
	"members/utils"
	"time"

	"github.com/rs/zerolog"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Sql struct {
		db     *bun.DB
		logger *zerolog.Logger
	}
	SqlInitializer = func(prov config.ConfigProvider) (*bun.DB, error)
)

var (
	_ Store = ((*Sql)(nil))
)

func NewSql(db *bun.DB) Store {
	return &Sql{db, nil}
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
	if _, err := sq.db.
		NewDropTable().
		Model(model).
		Cascade().
		Exec(context.TODO()); err != nil {
		// panic(err)
		sq.logger.Print(err)
	}
}

func (sq *Sql) CleanOldMembers(ctx context.Context, from time.Duration) error {
	if _, err := sq.db.NewDelete().
		Model((*common.Membership)(nil)).
		Where("last_health < ?", time.Now().Add(-from)).
		Exec(ctx); err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (sq *Sql) WithLogger(sub *zerolog.Logger) {
	sq.logger = sub
}

func (sq *Sql) QuoteCol(v string) string {
	return utils.QuoteCol(sq.db, v)
}

func (sq *Sql) Template(str string, vs ...any) string {
	return fmt.Sprintf(str, vs...)
}

func (sq *Sql) UpsertMembership(ctx context.Context, meta *common.Membership) error {
	if meta.Dns == "" {
		meta.Dns = meta.PublicAddress
	}

	base := sq.db.NewInsert().
		Model(meta)
	if _, err := base.
		Exec(context.TODO(), meta); err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (sq *Sql) GetMembers(ctx context.Context, kind ...common.Service) ([]*common.Membership, error) {
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
		sq.logger.Print(err)
		return nil, err
	}
	return out, nil
}

func (sq *Sql) CreateProject(ctx context.Context, project *common.ProtoProject, proto *common.ProtoMeta) error {
	project.Id = uuid.NewString()
	proto.Id = uuid.NewString()
	proto.Key = project.Id
	if _, err := sq.db.NewInsert().
		Model(project).
		Exec(ctx, project); err != nil && err != sql.ErrNoRows {
		sq.logger.Print(err)
		return err
	}
	if _, err := sq.db.NewInsert().
		Model(proto).
		Exec(ctx, proto); err != nil && err != sql.ErrNoRows {
		sq.logger.Print(err)
		return err
	}
	return nil
}

func (sq *Sql) Teardown() error {
	for _, tbl := range []interface{}{
		(*common.Membership)(nil),
		(*common.ProtoProject)(nil),
		(*common.RegisteredProto)(nil),
		(*common.ProtoMeta)(nil),
	} {
		drop(sq, tbl)
	}
	return nil
}

func (sq *Sql) create() error {
	if err := runCreate(sq, (*common.Membership)(nil)); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.ProtoProject)(nil)); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.ProtoMeta)(nil), func(ctq *bun.CreateTableQuery) *bun.CreateTableQuery {
		return ctq.ForeignKey(
			sq.Template(
				"(%s) references repos (%s) on update cascade on delete cascade",
				sq.QuoteCol("key"),
				sq.QuoteCol("id"),
			),
		)
	}); err != nil {
		return err
	}
	if err := runCreate(sq, (*common.RegisteredProto)(nil), func(ctq *bun.CreateTableQuery) *bun.CreateTableQuery {
		return ctq.ForeignKey(
			sq.Template(
				"(%s) references registration_meta (%s) on update cascade on delete cascade",
				sq.QuoteCol("id"),
				sq.QuoteCol("id"),
			))
	}); err != nil {
		return err
	}
	return nil
}

func (sq *Sql) Setup(cfg config.ConfigProvider) error {
	storg := cfg.GetConfig().Storage
	if storg.Drop {
		if err := sq.Teardown(); err != nil {
			return err
		}
	}
	if storg.Create {
		if err := sq.create(); err != nil {
			return err
		}
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
