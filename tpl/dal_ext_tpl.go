package tpl

var DalExtTpl = `package dal

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
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
`
