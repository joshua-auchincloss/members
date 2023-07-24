package utils

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Table(cols ...interface{}) table.Table {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgCyan).SprintfFunc()
	tbl := table.New(cols...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}
