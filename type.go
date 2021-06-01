package db2struct

import "strings"

func getType(col *ColDetail) string {

	// gorm soft delete
	if opts.GormTag && col.ColumnName == "deleted_at" && strings.ToUpper(col.DataType) == "TIMESTAMP" {
		return "gorm.DeletedAt"
	}

	switch strings.ToUpper(col.DataType) {
	case "TINYINT", "SMALLINT", "INT", "MEDIUMINT":
		return nullableType("int", col.IsNullable)
	case "BIGINT":
		return nullableType("int64", col.IsNullable)
	case "DECIMAL", "DOUBLE":
		return nullableType("float64", col.IsNullable)
	case "FLOAT":
		return nullableType("float32", col.IsNullable)
	case "BIT":
		return nullableType("uint64", col.IsNullable)
	case "BINARY", "BLOB", "LONGBLOB", "MEDIUMBLOB", "VARBINARY":
		return nullableType("[]byte", col.IsNullable)
	case "DATE", "DATETIME", "TIME", "TIMESTAMP":
		return nullableType("time.Time", col.IsNullable)
	default:
		return nullableType("string", col.IsNullable)
	}
}

func nullableType(t string, nullable bool) string {
	if nullable {
		return "*" + t
	}
	return t
}
