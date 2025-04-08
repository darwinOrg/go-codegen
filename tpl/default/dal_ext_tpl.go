package _default

import "strings"

var DalExtTpl = `package dal

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	"github.com/darwinOrg/go-common/page"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	{{if .HasDecimal}}"github.com/shopspring/decimal"{{end}}
	"time"
)

var {{.GoTable}}ExtDao = &{{.LowerCamelName}}ExtDao{}

type {{.LowerCamelName}}ExtDao struct {
}

type Query{{.GoTable}}Param struct {
    {{range .QueryColumns}}{{.GoName}} {{.ModelType}} ###json:"{{.LowerCamelName}}"
	{{- if ne .Comment ""}} remark:"{{.Comment}}"{{end}}###
	{{end}}
	*page.PageParam
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

func (d *{{.LowerCamelName}}ExtDao) Create(ctx *dgctx.DgContext, tc *daog.TransContext, {{.LowerCamelName}} *{{.GoTable}}) (int64, error) {
	{{$.LowerCamelName}}.CreatedAt = ttypes.NormalDatetime(time.Now())
	{{$.LowerCamelName}}.CreatedBy = ctx.UserId
	{{$.LowerCamelName}}.ModifiedAt = ttypes.NormalDatetime(time.Now())
	{{$.LowerCamelName}}.ModifiedBy = ctx.UserId

	_, err := {{.GoTable}}Dao.Insert(tc, {{.LowerCamelName}})
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Insert error: %v", err)
		return 0, dgerr.SYSTEM_ERROR
	}

	return {{.LowerCamelName}}.Id, nil
}

func (d *{{.LowerCamelName}}ExtDao) Modify(ctx *dgctx.DgContext, tc *daog.TransContext, {{.LowerCamelName}} *{{.GoTable}}) error {
	{{$.LowerCamelName}}.ModifiedAt = ttypes.NormalDatetime(time.Now())
	{{$.LowerCamelName}}.ModifiedBy = ctx.UserId

	_, err := {{.GoTable}}Dao.Update(tc, {{.LowerCamelName}})
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Update error: %v", err)
		return dgerr.SYSTEM_ERROR
	}

	return nil
}

func (d *{{.LowerCamelName}}ExtDao) DeleteById(ctx *dgctx.DgContext, tc *daog.TransContext, id int64) error {
	count, err := {{$.GoTable}}Dao.DeleteById(tc, id)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.DeleteById error: %v", err)
		return dgerr.SYSTEM_ERROR
	}
	if count == 0 {
		return dgerr.RECORD_NOT_EXISTS
	}

	return nil
}

func (d *{{.LowerCamelName}}ExtDao) Page(ctx *dgctx.DgContext, tc *daog.TransContext, param *Query{{.GoTable}}Param) (*page.PageList[{{.GoTable}}], error) {
	matcher := d.BuildMatcher(param)
	
	count, err := {{.GoTable}}Dao.Count(tc, matcher)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Count error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	if count == 0 {
		return page.EmptyPageList[{{.GoTable}}](param.PageNo, param.PageSize), nil
	}
	
	{{.LowerCamelName}}List, err := {{.GoTable}}Dao.QueryPageListMatcher(tc, matcher, daog.NewPager(param.PageSize, param.PageNo), daog.NewDescOrder({{$.GoTable}}Fields.CreatedAt))
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.QueryPageListMatcher error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	if len({{.LowerCamelName}}List) == 0 {
		return page.EmptyPageList[{{.GoTable}}](param.PageNo, param.PageSize), nil
	}
	
	return page.ListOf(param.PageNo, param.PageSize, int(count), {{.LowerCamelName}}List), nil
}

func (d *{{.LowerCamelName}}ExtDao) List(ctx *dgctx.DgContext, tc *daog.TransContext, param *Query{{.GoTable}}Param) ([]*{{.GoTable}}, error) {
	matcher := d.BuildMatcher(param)

	{{.LowerCamelName}}List, err := {{.GoTable}}Dao.QueryListMatcher(tc, matcher, daog.NewDescOrder({{$.GoTable}}Fields.CreatedAt))
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.QueryListMatcher error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}
	
	return {{.LowerCamelName}}List, nil
}

func (d *{{.LowerCamelName}}ExtDao) BuildMatcher(param *Query{{.GoTable}}Param) daog.Matcher {
	matcher := daog.NewMatcher()
	{{range .QueryColumns}}
	{{- if eq .DbType "string"}}
	if param.{{.GoName}} != "" {
		matcher.Eq({{$.GoTable}}Fields.{{.GoName}}, param.{{.GoName}})
	}
	{{- else if contains .DbType "int"}}
	if param.{{.GoName}} != 0 {
		matcher.Eq({{$.GoTable}}Fields.{{.GoName}}, param.{{.GoName}})
	}
	{{- end }}
	{{- end }}

	return matcher
}

func Get{{.GoTable}}Id({{.LowerCamelName}} *{{.GoTable}}) int64 {
	return {{.LowerCamelName}}.Id
}
`

func init() {
	DalExtTpl = strings.ReplaceAll(DalExtTpl, "###", "`")
}
