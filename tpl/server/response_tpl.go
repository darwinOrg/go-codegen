package _server

import "strings"

var ResponseTpl = `package model

{{if .HasDecimal}}import "github.com/shopspring/decimal"{{end}}

{{range .Responses}}
type {{.UpperCamelName}} struct {
	{{range .Models}}{{.FieldName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.DataType}} ###json:"{{.LowerCamelName}}"{{if .IsMediaUrl}} appendUid:"true"{{end}} {{if ne .EnumRemark ""}}title:"{{.EnumTitle}}" remark:"{{.EnumRemark}}"{{else}}remark:"{{.Remark}}"{{end}}###
	{{end}}
	{{if ne .ExtendName ""}}*{{.ExtendName}}{{- end}}
}
{{end}}
`

func init() {
	ResponseTpl = strings.ReplaceAll(ResponseTpl, "###", "`")
}
