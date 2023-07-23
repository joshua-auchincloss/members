package mem

import (
	"database/sql"
	"fmt"
	"log"
	"members/config"
	"members/utils"

	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var (
	fitempl = "file:%s?cache=shared"
)

func New(prov config.ConfigProvider) (*bun.DB, error) {
	fp := prov.GetConfig().Storage.URI
	if utils.ZeroStr(fp) {
		fp = fmt.Sprintf(fitempl, ":memory:")
	} else {
		fp = fmt.Sprintf(fitempl, fp)
	}
	log.Printf("using file: %s", fp)
	sqldb, err := sql.Open(sqliteshim.ShimName, fp)
	if err != nil {
		return nil, err
	}
	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}
