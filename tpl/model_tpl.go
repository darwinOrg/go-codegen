package tpl

var ModelTpl = `package model

import "github.com/darwinOrg/go-common/page"

type {{.GoTable}} struct {
    {{range .Columns}}{{.GoName}} {{.DbType}}
    {{end}}
}
`
