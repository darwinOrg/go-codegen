package _server

import "strings"

var RequestTpl = `package model

{{if .HasPage}}import "github.com/darwinOrg/go-common/page"{{end}}

{{range .Requests}}
type {{.UpperCamelName}} struct {
	{{range .Models}}{{.FieldName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.ModelType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}"{{if ne .VerifyRules ""}} binding:"{{.VerifyRules}}"{{end}} remark:"{{.Remark}}"###
	{{end}}
	{{if ne .ExtendName ""}}{{.ExtendName}}{{- end}}
	{{if .IsPage}}*page.PageParam{{- end}}
}
{{end}}
`

func init() {
	RequestTpl = strings.ReplaceAll(RequestTpl, "###", "`")
}
