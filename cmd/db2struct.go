package main

import (
	"context"
	"fmt"

	"github.com/gertd/go-pluralize"
	"github.com/himeteam/db2struct"
	"github.com/iGoogle-ink/gopher/errgroup"
	"github.com/iGoogle-ink/gopher/xlog"
)

func main() {
	db2struct.InitDBConn()
	defer db2struct.Close()
	opts := db2struct.GetOpts()
	var (
		eg  = errgroup.WithContext(context.Background())
		scs []*db2struct.StructCol
		tpl []byte
		err error
	)

	eg.Go(func(ctx context.Context) error {
		r := db2struct.GetTableDetail(opts.Table)
		scs = db2struct.ToStructCol(r)
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		tpl, err = db2struct.GetTpl(opts.Template)
		if err != nil {
			xlog.Errorf("db2struct.GetTpl(%s),err:%+v", opts.Template, err)
			return err
		}
		return nil
	})
	if err = eg.Wait(); err != nil {
		panic(err)
	}
	pluralizeClient := pluralize.NewClient()
	structStr := db2struct.GenerateStruct(string(tpl), pluralizeClient.Singular(opts.Table), scs)
	fmt.Println(structStr)
}
