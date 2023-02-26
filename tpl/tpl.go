package tpl

const (
	StructTpl = `package {{.Package}}

{{if .Imports}}import (
	{{range .Imports}}"{{.}}"{{end}}
){{end}}

{{if .Comment}}// {{.Name}} {{.Comment}}{{end}}
type {{.Name}} struct {
{{range .Fields}}   {{.Name}}  {{.Type}}  {{.FieldTag}}   // {{.Name}} {{.FieldName}} {{.Comment}}
{{end}}}


// TableName {{.Name}} {{.TableName}}
func ({{.Name}}) TableName() string {
	return "{{.TableName}}"
}`
)
