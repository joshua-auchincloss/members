package neo4j

import (
	"members/common"
	"members/storage/base"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GraphConvert(
	typ common.StorageType,
	conv ...func(neo4j.DriverWithContext) (base.BaseStore, error),
) GraphConverter {
	var do func(neo4j.DriverWithContext) (base.BaseStore, error)
	if len(conv) != 0 {
		do = conv[0]
	} else {
		do = base.NoOpBase(typ, NewNeo4jBase)
	}
	return func(d neo4j.DriverWithContext) (base.BaseStore, error) {
		if db, err := do(d); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}

}
