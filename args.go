package db2struct

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Opts struct {
	Host     string `short:"H" long:"host" description:"MySQL host" required:"false"`
	Port     string `long:"port" description:"MySQL post" required:"false"`
	User     string `short:"u" long:"user" description:"MySQL user" required:"true"`
	Password string `short:"p" long:"password" description:"MySQL password" required:"false"`
	Database string `short:"d" long:"database" description:"Database name" required:"true"`
	Table    string `short:"t" long:"table" description:"Table name" required:"true"`
	Template string `long:"template" description:"Template path" required:"false"`
	Package  string `long:"package" description:"Package name"`
	JsonTag  bool   `long:"json" description:"Add json tag"`
	GormTag  bool   `long:"gorm" description:"Add gorm tag"`
	XormTag  bool   `long:"xorm" description:"Add xorm tag"`
}

var opts Opts
var args []string

func init() {
	var err error
	args, err = flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(0)
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

	if opts.Package == "" {
		opts.Package = "model"
	}
}

func GetOpts() Opts {
	return opts
}

func GetArgs() []string {
	return args
}
