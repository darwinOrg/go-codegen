package _default

var HandlerTpl = `package handler

import (
	"{{.PackagePrefix}}/model"
	"{{.PackagePrefix}}/service"
	dgctx "github.com/darwinOrg/go-common/context"
	cm "github.com/darwinOrg/go-common/model"
	"github.com/darwinOrg/go-common/page"
	"github.com/darwinOrg/go-common/result"
	"github.com/gin-gonic/gin"
)

var {{.GoTable}}Handler = &{{.LowerCamelName}}Handler{}

type {{.LowerCamelName}}Handler struct {
}

func (h *{{.LowerCamelName}}Handler) Create(_ *gin.Context, ctx *dgctx.DgContext, req *model.Create{{.GoTable}}Req) *result.Result[int64] {
	{{.LowerCamelName}}Id, err := service.{{.GoTable}}Service.Create(ctx, req)
	if err != nil {
		return result.FailByError[int64](err)
	}
	return result.Success[int64]({{.LowerCamelName}}Id)
}

func (h *{{.LowerCamelName}}Handler) Modify(_ *gin.Context, ctx *dgctx.DgContext, req *model.Modify{{.GoTable}}Req) *result.Result[*result.Void] {
	err := service.{{.GoTable}}Service.Modify(ctx, req)
	if err != nil {
		return result.SimpleFailByError(err)
	}
	return result.SimpleSuccess()
}

func (h *{{.LowerCamelName}}Handler) Delete(_ *gin.Context, ctx *dgctx.DgContext, req *cm.IdReq) *result.Result[*result.Void] {
	err := service.{{.GoTable}}Service.Delete(ctx, req)
	if err != nil {
		return result.SimpleFailByError(err)
	}
	return result.SimpleSuccess()
}

func (h *{{.LowerCamelName}}Handler) Page(_ *gin.Context, ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) *result.Result[*page.PageList[model.{{.GoTable}}ListVo]] {
	pageList, err := service.{{.GoTable}}Service.Page(ctx, req)
	if err != nil {
		return result.FailByError[*page.PageList[model.{{.GoTable}}ListVo]](err)
	}
	return result.Success[*page.PageList[model.{{.GoTable}}ListVo]](pageList)
}

func (h *{{.LowerCamelName}}Handler) List(_ *gin.Context, ctx *dgctx.DgContext, req *model.Query{{.GoTable}}Req) *result.Result[[]*model.{{.GoTable}}ListVo] {
	list, err := service.{{.GoTable}}Service.List(ctx, req)
	if err != nil {
		return result.FailByError[[]*model.{{.GoTable}}ListVo](err)
	}
	return result.Success[[]*model.{{.GoTable}}ListVo](list)
}

func (h *{{.LowerCamelName}}Handler) Detail(_ *gin.Context, ctx *dgctx.DgContext, req *cm.IdReq) *result.Result[*model.{{.GoTable}}DetailResp] {
	detail, err := service.{{.GoTable}}Service.Detail(ctx, req.Id)
	if err != nil {
		return result.FailByError[*model.{{.GoTable}}DetailResp](err)
	}
	return result.Success[*model.{{.GoTable}}DetailResp](detail)
}

`
