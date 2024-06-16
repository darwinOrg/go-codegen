package _default

import "strings"

var ModelTpl = `package model

import (
	{{if .HasDaogType}}"github.com/rolandhe/daog/ttypes"{{end}}
	{{if .HasDecimal}}"github.com/shopspring/decimal"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
)

{{range .Requests}}
type {{.Name}} struct {
	{{range .Models}}
	{{.FieldName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.DataType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}" binding:"{{.VerifyRules}} remark:"{{.Remark}}"###
	{{end}}
	{{if ne .ExtendName ""}}{{.ExtendName}}{{end}}
	{{if .IsPage}}*page.PageParam{{end}}
}

{{range .Responses}}
type {{.Name}} struct {
	{{range .Models}}
	{{.FieldName}} {{if .IsArray}}[]{{end}}{{if .IsPointer}}*{{end}}{{.DataType}} ###json:"{{.LowerCamelName}}" {{if .IsMediaUrl}}appendUid:"true"{{end}} remark:"{{.Remark}}"###
	{{if ne .EnumModel ""}}{{.FieldName}}Name string ###json:"{{.LowerCamelName}}Name" title="EnumTitle" remark:"{{.EnumRemark}}"###{{end}}
	{{end}}
	{{if ne .ExtendName ""}}{{.ExtendName}}{{end}}
}
`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
