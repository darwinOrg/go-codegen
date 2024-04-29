package tpl

import "strings"

var ModelTpl = `package model

import (
	"github.com/darwinOrg/go-common/page"
	{{if eq .HasTime true}}"time"{{end}}
)

type Create{{.GoTable}}Req struct {
    {{range .CreateColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .IsNull false}} binding:"required"{{end}}
	{{- if eq .ModelType "time.Time"}} time_format:"2006-01-02 15:04:05"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type Modify{{.GoTable}}Req struct {
    {{range .ModifyColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .IsNull false}} binding:"required"{{end}}
	{{- if eq .ModelType "time.Time"}} time_format:"2006-01-02 15:04:05"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type Query{{.GoTable}}Req struct {
    {{range .CreateColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}" form:"{{.LowerCamelName}}"
	{{- if eq .ModelType "time.Time"}} time_format:"2006-01-02 15:04:05"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	*page.PageParam
}

type {{.GoTable}}ListVo struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .ModelType "time.Time"}} time_format:"2006-01-02 15:04:05"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

type {{.GoTable}}DetailVo struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if eq .ModelType "time.Time"}} time_format:"2006-01-02 15:04:05"{{end}}
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}

`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
