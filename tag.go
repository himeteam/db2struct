package db2struct

import (
	"fmt"
	"strings"
)

type getTagFunc func(cols *ColDetail) string

func getTags(col *ColDetail) []string {
	tags := make([]string, 0)
	var getTagFuncList []getTagFunc

	if opts.JsonTag {
		getTagFuncList = append(getTagFuncList, jsonTag)
	}
	if opts.GormTag {
		getTagFuncList = append(getTagFuncList, gormTag)
	}
	if opts.XormTag {
		getTagFuncList = append(getTagFuncList, xormTag)
	}

	for _, fn := range getTagFuncList {
		tag := fn(col)
		if tag == "" {
			continue
		}

		tags = append(tags, tag)
	}

	return tags
}

func jsonTag(col *ColDetail) string {
	return fmt.Sprintf("json:\"%s\"", col.ColumnName)
}

func gormTag(col *ColDetail) string {
	var tags []string
	tags = append(tags, "column:"+col.ColumnName)

	if strings.Contains(col.ColumnKey, "PRI") {
		tags = append(tags, "primaryKey")
	}

	if strings.Contains(col.Extra, "auto_increment") {
		tags = append(tags, "autoIncrement")
	}

	if col.ColumnName == "created_at" {
		tags = append(tags, "autoCreateTime")
	}

	if col.ColumnName == "updated_at" {
		tags = append(tags, "autoUpdateTime")
	}

	if !col.IsNullable {
		tags = append(tags, "not null")
	}

	if col.ColumnDefault != nil {
		tags = append(tags, "default: '"+*col.ColumnDefault+"'")
	}

	return "gorm:\"" + strings.Join(tags, ";") + "\""
}

func xormTag(col *ColDetail) string {
	tags := make([]string, 0)

	if strings.Contains(col.ColumnKey, "PRI") {
		tags = append(tags, "pk")
	}

	if strings.Contains(col.Extra, "auto_increment") {
		tags = append(tags, "autoincr")
	}

	if col.IsNullable {
		tags = append(tags, "null")
	} else {
		tags = append(tags, "notnull")
	}

	if col.ColumnDefault != nil {
		tags = append(tags, "default('"+*col.ColumnDefault+"')")
	}

	if col.DataType == "json" {
		tags = append(tags, "json")
	}

	if col.ColumnName == "created_at" {
		tags = append(tags, "created")
	}

	if col.ColumnName == "updated_at" {
		tags = append(tags, "updated")
	}

	if col.ColumnName == "deleted_at" {
		tags = append(tags, "deleted")
	}

	if len(tags) == 0 {
		return ""
	}

	return "xorm:\"" + strings.Join(tags, " ") + "\""
}
