package utils

import (
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
)

func QuoteCol(db *bun.DB, v string) string {
	quote := `"`
	switch db.Dialect().Name() {
	case dialect.MySQL:
		quote = "`"
	}
	return fmt.Sprintf("%s%s%s", quote, v, quote)
}
