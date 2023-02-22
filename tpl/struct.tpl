package {{.Package}}

{{if .ContainsTimeField}}import "time"{{end}}

{{if .Comment}}// {{.Name}} {{.Comment}}{{end}}
type {{.Name}} struct {
{{range .Fields}}   {{.Name}}  {{.FieldType}}  {{.FieldTag}}   // {{.Name}} {{.FieldName}} {{.Comment}}
{{end}}}


// TableName {{.Name}} {{.TableName}}
func ({{.Name}}) TableName() string {
    return {{.TableName}}
}