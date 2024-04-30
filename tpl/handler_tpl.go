package tpl

var HandlerExtTpl = `package handler

import (
	"{{.ProjectPath}}/model"
	"{{.ProjectPath}}/service"
	dgctx "github.com/darwinOrg/go-common/context"
	cm "github.com/darwinOrg/go-common/model"
	"github.com/darwinOrg/go-common/page"
	"github.com/darwinOrg/go-common/result"
)

var {{.GoTable}}Handler = &{{.LowerCamelName}}Handler{}

type {{.LowerCamelName}}Handler struct {
}

func (h *{{.LowerCamelName}}Handler) Create(ctx *dgctx.DgContext, req *model.Create{{.GoTable}}Req) *result.Result[int64] {
	{{.LowerCamelName}}Id, err := service.{{.GoTable}}Service.Create(ctx, req)
	if err != nil {
		return result.SimpleFail[int64](err.Error())
	}
	return result.Success[int64]({{.LowerCamelName}}Id)
}

func (h *{{.LowerCamelName}}Handler) Modify(ctx *dgctx.DgContext, req *model.Modify{{.GoTable}}Req) *result.Result[*result.Void] {
	err := service.{{.GoTable}}Service.Modify(ctx, req)
	if err != nil {
		return result.SimpleFailByError(err)
	}
	return result.SimpleSuccess()
}

func (h *{{.LowerCamelName}}Handler) DeleteById(ctx *dgctx.DgContext, req *cm.IdReq) *result.Result[*result.Void] {
	err := service.{{.GoTable}}Service.DeleteById(ctx, req.Id)
	if err != nil {
		return result.SimpleFailByError(err)
	}
	return result.SimpleSuccess()
}

func (h *{{.LowerCamelName}}Handler) Page(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) *result.Result[*page.PageList[model.{{.GoTable}}ListResp]] {
	pageList, err := service.{{.GoTable}}Service.Page(ctx, req)
	if err != nil {
		return result.SimpleFail[*page.PageList[model.{{.GoTable}}ListResp]](err.Error())
	}
	return result.Success[*page.PageList[model.{{.GoTable}}ListResp]](pageList)
}

func (h *{{.LowerCamelName}}Handler) List(ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) *result.Result[[]*model.{{.GoTable}}ListResp] {
	list, err := service.{{.GoTable}}Service.List(ctx, req)
	if err != nil {
		return result.SimpleFail[[]*model.{{.GoTable}}ListResp](err.Error())
	}
	return result.Success[[]*model.{{.GoTable}}ListResp](list)
}

func (h *{{.LowerCamelName}}Handler) Detail(ctx *dgctx.DgContext, req *cm.IdReq) *result.Result[*model.{{.GoTable}}DetailResp] {
	detail, err := service.{{.GoTable}}Service.Detail(ctx, req.Id)
	if err != nil {
		return result.SimpleFail[*model.{{.GoTable}}DetailResp](err.Error())
	}
	return result.Success[*model.{{.GoTable}}DetailResp](detail)
}

`
