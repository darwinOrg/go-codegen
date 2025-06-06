package _server

var EnumTpl = `package enum`

var EnumAppendTpl = `
{{range $index, $enum := .Enums}}
const (
	{{range $enum.Models}}{{.Code}} {{$enum.DataType}} = {{if eq $enum.DataType "string"}}"{{.Value}}"{{else}}{{.Value}}{{end}} // {{.Name}}
	{{end}}
)

var {{.UpperCamelName}}Map = map[{{$enum.DataType}}]string{
	{{range $enum.Models}}{{.Code}}: "{{.Name}}",
	{{end}}
}{{end}}
`

func init() {
	EnumTpl = EnumTpl + EnumAppendTpl
}
