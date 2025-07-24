package _default

import "strings"

var ModelTpl = `package model

import "github.com/darwinOrg/go-common/page"

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

type Query{{.GoTable}}Req struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	*page.PageParam
}

type {{.GoTable}}ListVo struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name string ###json:"{{.LowerCamelName}}Name"{{if ne .Comment ""}} title:"{{.EnumName}}名称" remark:"{{.EnumRemark}}"{{end}}###{{end}}
	{{end}}
}

type {{.GoTable}}DetailResp struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name string ###json:"{{.LowerCamelName}}Name"{{if ne .Comment ""}} title:"{{.EnumName}}名称" remark:"{{.EnumRemark}}"{{end}}###{{end}}
	{{end}}
}
`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
