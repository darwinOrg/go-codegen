package tpl

var DalExtTpl = `package dal

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/rolandhe/daog"
)

var {{.GoTable}}ExtDao = &{{.LowerCamelName}}ExtDao{}

type {{.LowerCamelName}}ExtDao struct {
}

func (d *{{.LowerCamelName}}ExtDao) GetById(ctx *dgctx.DgContext, tc *daog.TransContext, id int64) (*{{.GoTable}}, error) {
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

func (d *{{.LowerCamelName}}ExtDao) GetByIds(ctx *dgctx.DgContext, tc *daog.TransContext, ids []int64) ([]*{{.GoTable}}, error) {
	{{.LowerCamelName}}s, err := {{.GoTable}}Dao.GetByIds(tc, ids)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.GetByIds error: %v", err)
		return nil, dgerr.SYSTEM_ERROR
	}

	return {{.LowerCamelName}}s, nil
}

func (d *{{.LowerCamelName}}ExtDao) Insert(ctx *dgctx.DgContext, tc *daog.TransContext, ins *{{.GoTable}}) error {
	_, err := {{.GoTable}}Dao.Insert(tc, ins)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Insert error: %v", err)
		return dgerr.SYSTEM_ERROR
	}

	return nil
}

func (d *{{.LowerCamelName}}ExtDao) Update(ctx *dgctx.DgContext, tc *daog.TransContext, ins *{{.GoTable}}) error {
	_, err := {{.GoTable}}Dao.Update(tc, ins)
	if err != nil {
		dglogger.Errorf(ctx, "{{.GoTable}}Dao.Update error: %v", err)
		return dgerr.SYSTEM_ERROR
	}

	return nil
}

`
