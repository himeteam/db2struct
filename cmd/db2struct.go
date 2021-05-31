package main

import (
	"fmt"
	"github.com/himeteam/db2struct"
	"github.com/jessevdk/go-flags"
	"os"
)

var opts struct {
	Host     string `short:"H" long:"host" description:"MySQL host" required:"false"`
	Port     string `long:"port" description:"MySQL post" required:"false"`
	User     string `short:"u" long:"user" description:"MySQL user" required:"true"`
	Password string `short:"p" long:"password" description:"MySQL password" required:"false"`
	Database string `short:"d" long:"database" description:"Database name" required:"true"`
	Table    string `short:"t" long:"table" description:"Table name" required:"true"`
	Template string `long:"template" description:"Template path" required:"false"`
	JsonTag  bool   `long:"json" description:"Add json tag"`
	GormTag  bool   `long:"gorm" description:"Add gorm tag"`
	XormTag  bool   `long:"xorm" description:"Add xorm tag"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}

	if opts.Host == "" {
		opts.Host = "127.0.0.1"
	}

	if opts.Port == "" {
		opts.Port = "3306"
	}

	if opts.User == "" {
		opts.User = "root"
	}

	if opts.Template == "" {
		opts.Template = "templates/mysql.tpl"
	}

	tags := make(map[string]bool)
	if opts.JsonTag {
		tags["json"] = true
	}

	if opts.GormTag {
		tags["gorm"] = true
	}

	if opts.XormTag {
		tags["xorm"] = true
	}

	db2struct.InitDBConn(opts.Host, opts.Port, opts.User, opts.Password)
	defer db2struct.Close()

	r := db2struct.GetTableDetail(opts.Database, opts.Table)
	c := db2struct.ToStructCol(r, tags)

	tpl, _ := db2struct.GetTpl(opts.Template)
	structStr := db2struct.GenerateStruct(string(tpl), "task", c)
	fmt.Println(structStr)
}
