package tpl

import "strings"

var ModelTpl = `package model

import (
	"github.com/darwinOrg/go-common/page"
	{{if .ModelHasType}}"github.com/rolandhe/daog/ttypes"{{end}}
	{{if .ModelHasDecimal}}"github.com/shopspring/decimal"{{end}}
)

type Create{{.GoTable}}Req struct {
    {{range .CreateColumns}}{{.GoName}} {{.DbType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .IsNull false}} binding:"required"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type Modify{{.GoTable}}Req struct {
    {{range .ModifyColumns}}{{.GoName}} {{.DbType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .IsNull false}} binding:"required"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type Delete{{.GoTable}}Req struct {
    Id  int64 ###json:"id" binding:"required" remark:"id"###
}

type Query{{.GoTable}}Req struct {
    {{range .QueryColumns}}{{.GoName}} {{.DbType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	*page.PageParam
}

type {{.GoTable}}ListResp struct {
    {{range .QueryColumns}}{{.GoName}} {{.DbType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name string ###json:"{{.LowerCamelName}}Name"{{if ne .Comment ""}} title:"{{.EnumName}}名称" remark:"{{.EnumRemark}}"{{end}}###{{end}}
	{{end}}
}

type {{.GoTable}}DetailResp struct {
    {{range .QueryColumns}}{{.GoName}} {{.DbType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name string ###json:"{{.LowerCamelName}}Name"{{if ne .Comment ""}} title:"{{.EnumName}}名称" remark:"{{.EnumRemark}}"{{end}}###{{end}}
	{{end}}
}
`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
