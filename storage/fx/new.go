package storage_fx

import (
	"errors"
	"fmt"
	"members/common"
	"members/config"
	"members/logging"
	"members/storage"
	"members/storage/base"
	dgraph "members/storage/conns/graph/dgraph"
	"members/storage/conns/graph/neo4j"

	// "members/storage/conns/graph/neo4j"
	"members/storage/conns/sql"
	"members/storage/conns/sql/mem"
	"members/storage/conns/sql/mysql"
	pgres "members/storage/conns/sql/pgx"
	"members/utils"

	"github.com/dgraph-io/dgo"
	neoroot "github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type (
	Defn[
		DB any,
		T func(prov config.ConfigProvider) (DB, error),
		F func(DB) (base.BaseStore, error),
	] struct {
		Aliases []string
		Init    T
		Convert F
	}

	SqlDefn = Defn[
		*bun.DB,
		sql.SqlInitializer,
		sql.SqlConverter,
	]

	DGraphDefn = Defn[
		*dgo.Dgraph,
		dgraph.GraphInitializer,
		dgraph.GraphConverter,
	]

	Neo4jDefn = Defn[
		neoroot.DriverWithContext,
		neo4j.GraphInitializer,
		neo4j.GraphConverter,
	]
)

var (
	nosql   = errors.New("not sql")
	nograph = errors.New("not graph")

	Sql = []*SqlDefn{
		{Aliases: []string{
			"mem", "memory", "sqlite",
		}, Init: mem.New,
			Convert: sql.SqlConvert(common.StorageMemory),
		},
		{Aliases: []string{
			"postgres", "pg",
		}, Init: pgres.New,
			Convert: sql.SqlConvert(common.StoragePostgres),
		},
		{Aliases: []string{
			"mysql", "maria",
		}, Init: mysql.New,
			Convert: sql.SqlConvert(common.StorageMysql),
		},
	}

	DGraph = []*DGraphDefn{
		{Aliases: []string{
			"dgraph", "dgraph-mem", "graph-mem",
		}, Init: dgraph.New,
			Convert: dgraph.GraphConvert(common.StorageDGraph),
		},
	}

	Neo4j = []*Neo4jDefn{
		{Aliases: []string{
			"graph", "neo4j",
		}, Init: neo4j.New,
			Convert: neo4j.GraphConvert(common.StorageNeo4j),
		},
	}
)

func (d *Defn[T, Y, Z]) Is(s string) bool {
	return utils.AnyEq(d.Aliases, s)
}

func logger_for(prov config.ConfigProvider, root *zerolog.Logger) *zerolog.Logger {
	store := prov.GetConfig().Storage
	sub := logging.WithSub(root, "storage", func(ctx zerolog.Context) zerolog.Context {
		ctx = ctx.Str("storage-kind", store.Kind).
			Str("storage-uri", store.URI).
			Bool("storage-debug", store.Debug)
		if store.Debug {
			ctx = ctx.
				Uint32("storage-port", store.Port).
				Bool("storage-create-on-init", store.Create).
				Bool("storage-drop-on-init", store.Drop)
		}
		return ctx
	})
	return sub
}

func New[
	DB any,
	T func(prov config.ConfigProvider) (DB, error),
	F func(DB) (base.BaseStore, error),
](
	prov config.ConfigProvider,
	root *zerolog.Logger,
	defns []*Defn[DB, T, F],
) (base.StoreFactory, error) {
	st := prov.GetConfig().Storage.Kind
	for _, defn := range defns {
		if defn.Is(st) {
			b := storage.NewStoreFactory(
				prov,
				root,
				defn.Init,
				defn.Convert,
			)
			return b, nil
		}
	}
	return nil, fmt.Errorf("invalid storage type: %s", st)
}
