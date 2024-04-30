package tpl

var ServiceExtTpl = `package service

import (
	"{{.ProjectPath}}/converter"
	"{{.ProjectPath}}/dal"
	"{{.ProjectPath}}/model"
	dgctx "github.com/darwinOrg/go-common/context"
	daogext "github.com/darwinOrg/daog-ext"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	"github.com/darwinOrg/go-common/page"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/rolandhe/daog"
)

var {{.GoTable}}Service = &{{.LowerCamelName}}Service{}

type {{.LowerCamelName}}Service struct {
}

func (s *{{.LowerCamelName}}Service) Create(ctx *dgctx.DgContext, req *model.Create{{.GoTable}}Req) (int64, error) {
	{{.LowerCamelName}} := converter.{{.GoTable}}Converter.CreateModel2Entity(req)

	return 0, nil
}

`
