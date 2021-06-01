package {{.Package}}

type {{.TableName}} struct {
    {{range .Cols}} {{.Name}} {{.Type}} {{if ne .Tag "" }}{{.Tag | raw}}{{end}} {{if ne .Comment ""}}// {{.Comment}}{{end}}
{{end}} }