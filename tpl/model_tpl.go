package tpl

import "strings"

var ModelTpl = `package model

import "github.com/darwinOrg/go-common/page"

type Create{{.GoTable}}Req struct {
    {{range .CreateColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.JsonName}}"{{if eq .IsNull false}} binding:"required"{{end}}{{if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
}
`

func init() {
	ModelTpl = strings.ReplaceAll(ModelTpl, "###", "`")
}
