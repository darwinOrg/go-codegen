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

type {{.LowerCamelName}}Service struct{}

func (s *{{.LowerCamelName}}Service) Create(ctx *dgctx.DgContext, req *model.Create{{.GoTable}}Req) (int64, error) {
	{{.LowerCamelName}} := converter.{{.GoTable}}Converter.CreateReq2Entity(req)

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

		converter.{{.GoTable}}Converter.FillEntityWithModifyReq({{.LowerCamelName}}, req)

		return dal.{{.GoTable}}ExtDao.Modify(ctx, tc, {{.LowerCamelName}})
	})
}

func (s *{{.LowerCamelName}}Service) Delete(ctx *dgctx.DgContext, req *cm.IdReq) error {
	return daogext.Write(ctx, func(tc *daog.TransContext) error {
		return dal.{{.GoTable}}ExtDao.DeleteById(ctx, tc, req.Id)
	})
}

func (s *{{.LowerCamelName}}Service) Page(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) (*page.PageList[model.{{.GoTable}}ListVo], error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) (*page.PageList[model.{{.GoTable}}ListVo], error) {
		queryParam := converter.{{.GoTable}}Converter.QueryReq2Param(req)
		pl, err := dal.{{.GoTable}}ExtDao.Page(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listVos := dgcoll.MapToList(pl.List, converter.{{.GoTable}}Converter.Entity2ListVo)

		return page.ListOf[model.{{.GoTable}}ListVo](pl.PageNo, pl.PageSize, pl.TotalCount, listVos), nil
	})
}

func (s *{{.LowerCamelName}}Service) List(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) ([]*model.{{.GoTable}}ListVo, error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) ([]*model.{{.GoTable}}ListVo, error) {
		queryParam := converter.{{.GoTable}}Converter.QueryReq2Param(req)
		{{.LowerCamelName}}List, err := dal.{{.GoTable}}ExtDao.List(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listVos := dgcoll.MapToList({{.LowerCamelName}}List, converter.{{.GoTable}}Converter.Entity2ListVo)

		return listVos, nil
	})
}

func (s *{{.LowerCamelName}}Service) Detail(ctx *dgctx.DgContext, id int64) (*model.{{.GoTable}}DetailResp, error) {
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) (*model.{{.GoTable}}DetailResp, error) {
		{{.LowerCamelName}}, err := dal.{{.GoTable}}ExtDao.MustGetById(ctx, tc, id)
		if err != nil {
			return nil, err
		}

		detailResp := converter.{{.GoTable}}Converter.Entity2DetailResp({{.LowerCamelName}})

		return detailResp, nil
	})
}

`
