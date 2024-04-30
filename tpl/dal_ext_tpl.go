package tpl

var DalExtTpl = `package dal

import (
	"{{.ProjectPath}}/model"
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	"github.com/darwinOrg/go-common/page"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/rolandhe/daog"
)

var {{.GoTable}}ExtDao = &{{.LowerCamelName}}ExtDao{}

type {{.LowerCamelName}}ExtDao struct {
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
	matcher := d.buildMatcher(req)
	
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
	matcher := d.buildMatcher(req)

	{{.LowerCamelName}}List, err := {{.GoTable}}Dao.QueryListMatcher(tc, matcher, daog.NewDescOrder({{$.GoTable}}Fields.CreatedAt))
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.QueryListMatcher error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	
	return {{.LowerCamelName}}List, nil
}

func (d *{{.LowerCamelName}}ExtDao) buildMatcher(req *model.Query{{.GoTable}}Req) daog.Matcher {
	matcher := daog.NewMatcher()
	{{range $index, $column := .QueryColumns -}}
	{{- if eq $column.ModelType "string"}}
	if req.{{$column.GoName}} != "" {
		matcher.Eq({{$.GoTable}}Fields.{{$column.GoName}}, req.{{$column.GoName}})
	}
	{{- else if contains $column.ModelType "int"}}
	if req.{{$column.GoName}} != 0 {
		matcher.Eq({{$.GoTable}}Fields.{{$column.GoName}}, req.{{$column.GoName}})
	}
	{{- end }}
	{{- end }}

	return matcher
}

`
