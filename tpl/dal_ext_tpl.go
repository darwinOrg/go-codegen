package tpl

var DalExtTpl = `package dal

import (
	"{{.ProjectPath}}/model"
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	"github.com/darwinOrg/go-common/page"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"time"
)

var {{.GoTable}}ExtDao = &{{.LowerCamelName}}ExtDao{}

type {{.LowerCamelName}}ExtDao struct {
}

func (d *{{.LowerCamelName}}ExtDao) Create(ctx *dgctx.DgContext, tc *daog.TransContext, {{.LowerCamelName}} *{{.GoTable}}) (int64, error) {
	now := ttypes.NormalDatetime(time.Now())
	{{.LowerCamelName}}.CreatedAt = now
	{{.LowerCamelName}}.ModifiedAt = now
	{{.LowerCamelName}}.CreatedBy = ctx.UserId
	{{.LowerCamelName}}.ModifiedBy = ctx.UserId

	_, err := {{.GoTable}}Dao.Insert(tc, {{.LowerCamelName}})
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Insert error: %v", err)
		return 0, dgerr.SYSTEM_ERROR
	}

	return {{.LowerCamelName}}.Id, nil
}

func (d *{{.LowerCamelName}}ExtDao) Modify(ctx *dgctx.DgContext, tc *daog.TransContext, {{.LowerCamelName}} *{{.GoTable}}) error {
	{{.LowerCamelName}}.ModifiedAt = ttypes.NormalDatetime(time.Now())
	{{.LowerCamelName}}.ModifiedBy = ctx.UserId

	_, err := {{.GoTable}}Dao.Update(tc, {{.LowerCamelName}})
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Update error: %v", err)
		return dgerr.SYSTEM_ERROR
	}

	return nil
}

func (d *{{.LowerCamelName}}ExtDao) MustGetById(ctx *dgctx.DgContext, tc *daog.TransContext, id int64) (*{{.GoTable}}, error) {
	{{.LowerCamelName}}, err := {{.GoTable}}Dao.GetById(tc, id)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.GetById error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	if {{.LowerCamelName}} == nil {
		return nil, dgerr.RECORD_NOT_EXISTS
	}

	return {{.LowerCamelName}}, nil
}

func (d *{{.LowerCamelName}}ExtDao) Page(ctx *dgctx.DgContext, tc *daog.TransContext, req *model.Query{{.GoTable}}Req) (*page.PageList[{{.GoTable}}], error) {
	matcher := d.BuildMatcher(req)
	
	count, err := {{.GoTable}}Dao.Count(tc, matcher)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Count error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	if count == 0 {
		return page.EmptyPageList[{{.GoTable}}](req.PageNo, req.PageSize), nil
	}
	
	{{.LowerCamelName}}List, err := {{.GoTable}}Dao.QueryPageListMatcher(tc, matcher, daog.NewPager(req.PageSize, req.PageNo), daog.NewDescOrder({{$.GoTable}}Fields.CreatedAt))
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.QueryPageListMatcher error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	if len({{.LowerCamelName}}List) == 0 {
		return page.EmptyPageList[{{.GoTable}}](req.PageNo, req.PageSize), nil
	}
	
	return page.ListOf(req.PageNo, req.PageSize, int(count), {{.LowerCamelName}}List), nil
}

func (d *{{.LowerCamelName}}ExtDao) List(ctx *dgctx.DgContext, tc *daog.TransContext, req *model.Query{{.GoTable}}Req) ([]*{{.GoTable}}, error) {
	matcher := d.BuildMatcher(req)

	{{.LowerCamelName}}List, err := {{.GoTable}}Dao.QueryListMatcher(tc, matcher, daog.NewDescOrder({{$.GoTable}}Fields.CreatedAt))
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.QueryListMatcher error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	
	return {{.LowerCamelName}}List, nil
}

func (d *{{.LowerCamelName}}ExtDao) BuildMatcher(req *model.Query{{.GoTable}}Req) daog.Matcher {
	matcher := daog.NewMatcher()
	{{range .QueryColumns}}
	{{- if eq .ModelType "string"}}
	if req.{{.GoName}} != "" {
		matcher.Eq({{$.GoTable}}Fields.{{.GoName}}, req.{{.GoName}})
	}
	{{- else if contains .ModelType "int"}}
	if req.{{.GoName}} != 0 {
		matcher.Eq({{$.GoTable}}Fields.{{.GoName}}, req.{{.GoName}})
	}
	{{- end }}
	{{- end }}

	return matcher
}

`
