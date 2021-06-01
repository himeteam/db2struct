package main

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/himeteam/db2struct"
)

func main() {
	db2struct.InitDBConn()
	defer db2struct.Close()

	opts := db2struct.GetOpts()
	r := db2struct.GetTableDetail(opts.Table)
	c := db2struct.ToStructCol(r)

	tpl, _ := db2struct.GetTpl(opts.Template)
	pluralizeClient := pluralize.NewClient()
	structStr := db2struct.GenerateStruct(string(tpl), pluralizeClient.Singular(opts.Table), c)
	fmt.Println(structStr)
}
