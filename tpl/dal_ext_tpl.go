package tpl

var DalExtTpl = `package dal

import (
	"github.com/darwinOrg/go-common/page"
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dglogger "github.com/darwinOrg/go-logger"
	"{{.ProjectPath}}/model"
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

func (d *{{.LowerCamelName}}ExtDao) Query(ctx *dgctx.DgContext, tc *daog.TransContext, req *model.Query{{.GoTable}}Req) (*page.PageList[{{.GoTable}}], error) {
	matcher := daog.NewMatcher()
	{{range $index, $column := .QueryColumns -}}
	{{- if eq $column.ModelType "string"}}
	if req.{{$column.GoName}} != "" {
		matcher.Eq({{$.GoTable}}Fields.{{$column.GoName}}, req.{{$column.GoName}})
	}
	{{- end }}
	{{- end }}

	return nil, nil
}

`
