package tpl

var EnumTpl = `package enum

{{range $column, $enums := .EnumMap}}
type {{$.GoTable}}{{$column.GoName}} {{$column.DbType}}
var (
	{{range $index, $enum := $enums}}
	{{$.GoTable}}{{$column.GoName}}_{{$enum.Key}} {{$.GoTable}}{{$column.GoName}} = {{if eq $column.DbType "string"}}"{{$enum.Key}}"{{else}}{{$enum.Key}}{{end}} // {{$enum.Value}}
	{{end}}
)
{{end}}
`
