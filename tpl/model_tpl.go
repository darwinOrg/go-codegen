package tpl

import "strings"

var ModelTpl = `package model

import (
	"github.com/darwinOrg/go-common/page"
	{{if .HasDecimal}}"github.com/shopspring/decimal"{{end}}
	{{if .HasType}}"github.com/rolandhe/daog/ttypes"{{end}}
)

type Create{{.GoTable}}Req struct {
    {{range .CreateColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .IsNull false}} binding:"required"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type Modify{{.GoTable}}Req struct {
    {{range .ModifyColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .IsNull false}} binding:"required"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type Delete{{.GoTable}}Req struct {
    Id  int64 ###json:"id" binding:"required" remark:"id"###
}

type Query{{.GoTable}}Req struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	*page.PageParam
}

type {{.GoTable}}ListResp struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type {{.GoTable}}DetailResp struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}
`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
