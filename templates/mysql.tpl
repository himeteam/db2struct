package {{.Package}}

type {{.TableName}} struct {
    {{range .Cols}} {{.FieldName}} {{.Type}} {{if ne .Tag "" }}{{.Tag | raw}}{{end}} {{if ne .Comment ""}}   // {{.Comment}}{{end}}
    {{end}} }

func (m *{{.TableName}}) TableName() string {
	return "{{.DBTableName}}"
}
