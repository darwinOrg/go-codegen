package tpl

var RouterExtTpl = `package router

import (
	"{{.ProjectPath}}/handler"
	"{{.ProjectPath}}/model"
	cm "github.com/darwinOrg/go-common/model"
	"github.com/darwinOrg/go-common/page"
	"github.com/darwinOrg/go-common/result"
	"github.com/darwinOrg/go-web/wrapper"
	"github.com/gin-gonic/gin"
)

func Bind{{.GoTable}}Router(rg *gin.RouterGroup) {
	g := rg.Group("/api/{{.KebabName}}/v1")
	
	wrapper.Post(&wrapper.RequestHolder[model.Create{{.GoTable}}Req, *result.Result[int64]]{
		Remark:       "{{.TableComment}} - 创建",
		RouterGroup:  g,
		RelativePath: "create",
		BizHandler:   handler.{{.GoTable}}Handler.Create,
	})

	wrapper.Post(&wrapper.RequestHolder[model.Modify{{.GoTable}}Req, *result.Result[*result.Void]]{
		Remark:       "{{.TableComment}} - 修改",
		RouterGroup:  g,
		RelativePath: "modify",
		BizHandler:   handler.{{.GoTable}}Handler.Modify,
	})

	wrapper.Post(&wrapper.RequestHolder[cm.IdReq, *result.Result[*result.Void]]{
		Remark:       "{{.TableComment}} - 删除",
		RouterGroup:  g,
		RelativePath: "delete",
		BizHandler:   handler.{{.GoTable}}Handler.DeleteById,
	})

	wrapper.Post(&wrapper.RequestHolder[model.Query{{.GoTable}}Req, *result.Result[*page.PageList[model.{{.GoTable}}ListResp]]]{
		Remark:       "{{.TableComment}} - 分页查询",
		RouterGroup:  g,
		RelativePath: "page",
		BizHandler:   handler.{{.GoTable}}Handler.Page,
	})

	wrapper.Post(&wrapper.RequestHolder[model.Query{{.GoTable}}Req, *result.Result[[]*model.{{.GoTable}}ListResp]]{
		Remark:       "{{.TableComment}} - 列表查询",
		RouterGroup:  g,
		RelativePath: "list",
		BizHandler:   handler.{{.GoTable}}Handler.List,
	})

	wrapper.Post(&wrapper.RequestHolder[cm.IdReq, *result.Result[*model.{{.GoTable}}DetailResp]]{
		Remark:       "{{.TableComment}} - 详情",
		RouterGroup:  g,
		RelativePath: "detail",
		BizHandler:   handler.{{.GoTable}}Handler.Detail,
	})
}
`
