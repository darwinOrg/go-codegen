package _default

var EnumTpl = `package enum

{{range $column, $enums := .EnumMap}}
const (
	{{range $index, $enum := $enums}}{{$column.GoName}}_{{$enum.Key}} {{$column.DbType}} = {{if eq $column.DbType "string"}}"{{$enum.Key}}"{{else}}{{$enum.Key}}{{end}} // {{$enum.Value}}
	{{end}}
)
var {{$column.GoName}}Map = map[{{$column.DbType}}]string{
	{{range $index, $enum := $enums}}{{$column.GoName}}_{{$enum.Key}}: "{{$enum.Value}}",
	{{end}}
}
{{end}}
`
