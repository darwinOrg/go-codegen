package _default

var ServiceTpl = `package service

import (
	daogext "github.com/darwinOrg/daog-ext"
	"{{.PackagePrefix}}/converter"
	"{{.PackagePrefix}}/dal"
	"{{.PackagePrefix}}/model"
	cm "github.com/darwinOrg/go-common/model"
	dgcoll "github.com/darwinOrg/go-common/collection"
	dgctx "github.com/darwinOrg/go-common/context"
	"github.com/darwinOrg/go-common/page"
	"github.com/rolandhe/daog"
)

var {{.GoTable}}Service = &{{.LowerCamelName}}Service{}

type {{.LowerCamelName}}Service struct {
}

func (s *{{.LowerCamelName}}Service) Create(ctx *dgctx.DgContext, req *model.Create{{.GoTable}}Req) error {
	{{.LowerCamelName}} := converter.{{.GoTable}}Converter.CreateModel2Entity(req)

	return daogext.Write(ctx, func(tc *daog.TransContext) error {
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

func (s *{{.LowerCamelName}}Service) Delete(ctx *dgctx.DgContext, req *cm.IdReq) error {
	return daogext.Write(ctx, func(tc *daog.TransContext) error {
		return dal.{{.GoTable}}ExtDao.DeleteById(ctx, tc, req.Id)
	})
}

func (s *{{.LowerCamelName}}Service) Page(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) (*page.PageList[model.{{.GoTable}}ListResp], error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) (*page.PageList[model.{{.GoTable}}ListResp], error) {
		queryParam := converter.{{.GoTable}}Converter.QueryModel2Param(req)
		pl, err := dal.{{.GoTable}}ExtDao.Page(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList(pl.List, converter.{{.GoTable}}Converter.Entity2ListModel)

		return page.ListOf[model.{{.GoTable}}ListResp](pl.PageNo, pl.PageSize, pl.TotalCount, listModels), nil
	})
}

func (s *{{.LowerCamelName}}Service) List(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) ([]*model.{{.GoTable}}ListResp, error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) ([]*model.{{.GoTable}}ListResp, error) {
		queryParam := converter.{{.GoTable}}Converter.QueryModel2Param(req)
		{{.LowerCamelName}}List, err := dal.{{.GoTable}}ExtDao.List(ctx, tc, queryParam)
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
