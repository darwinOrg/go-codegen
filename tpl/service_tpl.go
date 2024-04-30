package tpl

var ServiceExtTpl = `package service

import (
	daogext "github.com/darwinOrg/daog-ext"
	"{{.ProjectPath}}/converter"
	"{{.ProjectPath}}/dal"
	"{{.ProjectPath}}/model"
	dgctx "github.com/darwinOrg/go-common/context"
	"github.com/darwinOrg/go-common/page"
	"github.com/rolandhe/daog"
)

var {{.GoTable}}Service = &{{.LowerCamelName}}Service{}

type {{.LowerCamelName}}Service struct {
}

func (s *{{.LowerCamelName}}Service) Create(ctx *dgctx.DgContext, req *model.Create{{.GoTable}}Req) (int64, error) {
	{{.LowerCamelName}} := converter.{{.GoTable}}Converter.CreateModel2Entity(req)

	return daogext.WriteWithResult(ctx, func(tc *daog.TransContext) (int64, error) {
		return dal.{{.GoTable}}ExtDao.Create(ctx, tc, {{.LowerCamelName}})
	})
}

func (s *{{.LowerCamelName}}Service) Modify(ctx *dgctx.DgContext, req *model.Modify{{.GoTable}}Req) error {
	return daogext.Write(ctx, func(tc *daog.TransContext) error {
		{{.LowerCamelName}}, err := dal.{{.GoTable}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return err
		}

		converter.{{.GoTable}}Converter.FillEntityWithModifyModel({{.LowerCamelName}}, req)

		return dal.{{.GoTable}}ExtDao.Modify(ctx, tc, {{.LowerCamelName}})
	})
}

`
