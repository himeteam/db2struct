package db2struct

import (
	"bytes"
	"embed"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

//go:embed templates
var tpls embed.FS

func raw(x string) interface{} {
	return template.HTML(x)
}

func GenerateStruct(tpl string, table string, cols []*StructCol) string {
	t := template.New(table)
	t = t.Funcs(template.FuncMap{"raw": raw})
	t, err := t.Parse(tpl)
	if err != nil {
		panic(err)
	}

	var data struct {
		Package     string
		TableName   string
		DBTableName string
		Cols        []*StructCol
	}

	data.Package = opts.Package
	data.TableName = strings.Title(strcase.ToCamel(table))
	data.DBTableName = table
	data.Cols = cols

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, data)

	if err != nil {
		panic(err)
	}

	o, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(o)
}

func GetTpl(fileName string) (tplContent []byte, err error) {
	if _, err = os.Stat(fileName); err == nil {
		return ioutil.ReadFile(fileName)
	}

	pwd, _ := os.Getwd()
	f := filepath.Join(pwd, fileName)

	if _, err = os.Stat(f); err == nil {
		return ioutil.ReadFile(f)
	}

	file, err := tpls.Open(fileName)
	if err != nil {
		return
	}

	return ioutil.ReadAll(file)
}
