package tpl

var ServiceExtTpl = `package service

import (
	daogext "github.com/darwinOrg/daog-ext"
	"{{.ProjectPath}}/converter"
	"{{.ProjectPath}}/dal"
	"{{.ProjectPath}}/model"
	dgcoll "github.com/darwinOrg/go-common/collection"
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

func (s *{{.LowerCamelName}}Service) Page(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) (*page.PageList[model.{{.GoTable}}ListResp], error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) (*page.PageList[model.{{.GoTable}}ListResp], error) {
		pl, err := dal.{{.GoTable}}ExtDao.Page(ctx, tc, req)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList(pl.List, converter.{{.GoTable}}Converter.Entity2ListModel)

		return page.ListOf[model.{{.GoTable}}ListResp](pl.PageNo, pl.PageSize, pl.TotalCount, listModels), nil
	})
}

func (s *{{.LowerCamelName}}Service) List(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) ([]*model.{{.GoTable}}ListResp, error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) ([]*model.{{.GoTable}}ListResp, error) {
		{{.LowerCamelName}}List, err := dal.{{.GoTable}}ExtDao.List(ctx, tc, req)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList({{.LowerCamelName}}List, converter.{{.GoTable}}Converter.Entity2ListModel)

		return listModels, nil
	})
}

func (s *{{.LowerCamelName}}Service) Detail(ctx *dgctx.DgContext, id int64) (*model.{{.GoTable}}DetailResp, error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) (*model.{{.GoTable}}DetailResp, error) {
		{{.LowerCamelName}}, err := dal.{{.GoTable}}ExtDao.MustGetById(ctx, tc, id)
		if err != nil {
			return nil, err
		}

		detailModel := converter.{{.GoTable}}Converter.Entity2DetailModel({{.LowerCamelName}})

		return detailModel, nil
	})
}

`
