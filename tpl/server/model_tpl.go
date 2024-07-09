package _server

import "strings"

var ModelTpl = `package model

{{if .HasDecimal}}import "github.com/shopspring/decimal"{{end}}
{{if .HasPage}}import "github.com/darwinOrg/go-common/page"{{end}}

{{range .Requests}}
type {{.UpperCamelName}} struct {
	{{range .Models}}{{.UpperCamelName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.DataType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}"{{if ne .VerifyRules ""}} binding:"{{.VerifyRules}}"{{end}} remark:"{{.Remark}}"###
	{{end}}
	{{if ne .ExtendName ""}}{{.ExtendName}}{{- end}}
	{{if .IsPage}}*page.PageParam{{- end}}
}
{{end}}

{{range .Responses}}
type {{.UpperCamelName}} struct {
	{{range .Models}}{{.UpperCamelName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.DataType}} ###json:"{{.LowerCamelName}}"{{if .IsMediaUrl}} appendUid:"true"{{end}} {{if ne .EnumRemark ""}}title:"{{.EnumTitle}}" remark:"{{.EnumRemark}}"{{else}}remark:"{{.Remark}}"{{end}}###
	{{end}}
	{{if ne .ExtendName ""}}*{{.ExtendName}}{{- end}}
}
{{end}}
`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
