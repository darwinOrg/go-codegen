package _server

import "strings"

var ResponseTpl = `package model

{{range .Responses}}
type {{.UpperCamelName}} struct {
	{{range .Models}}{{.FieldName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.ModelType}} ###json:"{{.LowerCamelName}}"{{if .IsMediaUrl}} appendUid:"true"{{end}} {{if ne .EnumRemark ""}}title:"{{.EnumTitle}}" remark:"{{.EnumRemark}}"{{else}}remark:"{{.Remark}}"{{end}}###
	{{end}}
	{{if ne .ExtendName ""}}*{{.ExtendName}}{{- end}}
}
{{end}}
`

func init() {
	ResponseTpl = strings.ReplaceAll(ResponseTpl, "###", "`")
}
