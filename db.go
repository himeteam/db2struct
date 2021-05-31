package db2struct

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"net/url"
	"strings"
)

type ColDetail struct {
	ColumnName    string
	DataType      string
	ColumnType    string
	ColumnKey     string
	Extra         string
	ColumnComment string
	IsNullable    bool
	ColumnDefault *string
}

var db *sql.DB

func GetDSN(host, port, user, passwd string) string {
	connectConfigs := make(url.Values)
	connectConfigs.Set("charset", "utf8mb4")
	connectConfigs.Set("parseTime", "True")
	connectConfigs.Set("loc", "Local")

	if port == "" {
		port = "3306"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?%s", user, passwd, host, port, connectConfigs.Encode())
}

func InitDBConn(host, port, user, passwd string) {
	var err error
	db, err = sql.Open("mysql", GetDSN(host, port, user, passwd))
	if err != nil {
		panic(err)
	}

}

func Close() {
	if db == nil {
		return
	}
	db.Close()
}

func GetTableDetail(dbName, table string) []ColDetail {
	rows, err := db.Query("select "+
		"COLUMN_NAME,"+
		"DATA_TYPE,"+
		"COLUMN_TYPE,"+
		"COLUMN_KEY,"+
		"EXTRA,"+
		"COLUMN_COMMENT,"+
		"IS_NULLABLE,"+
		"COLUMN_DEFAULT "+
		"from information_schema.columns where table_schema = ? and table_name = ? order by ORDINAL_POSITION ASC", dbName, table)

	if err != nil {
		panic(err)
	}

	cols := make([]ColDetail, 0)

	for rows.Next() {
		col := ColDetail{}
		var nullable string
		err = rows.Scan(
			&col.ColumnName,
			&col.DataType,
			&col.ColumnType,
			&col.ColumnKey,
			&col.Extra,
			&col.ColumnComment,
			&nullable,
			&col.ColumnDefault,
		)

		if err != nil {
			panic(err)
		}

		if strings.ToUpper(nullable) == "YES" {
			col.IsNullable = true
		}

		cols = append(cols, col)
	}
	return cols
}

type StructCol struct {
	Name       string
	ColName    string
	Type       string
	Comment    string
	IsNullable bool
	IsPri      bool
	Tag        string
}

func ToStructCol(cols []ColDetail, enableTags map[string]bool) []StructCol {
	sc := make([]StructCol, 0)
	for _, col := range cols {
		c := StructCol{}
		c.Name = strings.Title(strcase.ToCamel(col.ColumnName))
		c.ColName = col.ColumnName
		c.IsNullable = col.IsNullable
		c.Comment = strings.ReplaceAll(col.ColumnComment, "\n", " ")
		c.IsPri = strings.Contains(col.Extra, "PRI")
		c.Type = getType(col)
		tags := getTags(col, enableTags)

		if len(tags) > 0 {
			c.Tag = "`" + strings.Join(tags, " ") + "`"
		}
		sc = append(sc, c)
	}
	return sc
}
