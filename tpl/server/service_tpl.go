package _server

var ServiceTpl = `package service

import (
	dgctx "github.com/darwinOrg/go-common/context"
	{{if .HasModel}}"{{.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	{{if .HasDbTable}}
		{{if .HasModel}}"{{.PackagePrefix}}/converter"{{end}}
		{{if .HasQuery}}dgcoll "github.com/darwinOrg/go-common/collection"{{end}}
		"{{.PackagePrefix}}/dal"
		daogext "github.com/darwinOrg/daog-ext"
		"github.com/rolandhe/daog"
	{{end}}
)

var {{.GroupUpperCamel}}Service = &{{.GroupLowerCamel}}Service{}
type {{.GroupLowerCamel}}Service struct {
}
`

var ServiceAppendTpl = `
{{range .Models}}
func (s *{{$.GroupLowerCamel}}Service) {{.MethodNameExp}}(ctx *dgctx.DgContext, req *{{.RequestModelNameExp}}) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}(int64, error){{end}} {
	{{- if eq .DbModelName ""}}
	// TODO
	return {{if ne .ResponseModelName ""}}nil, {{end}}nil
	{{- else if eq .InterfaceType "新建"}}
	return daogext.WriteWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}(int64, error){{end}} {
		{{.DbTableLowerCamel}} := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Entity(req)
		return dal.{{.DbTableUpperCamel}}ExtDao.Create(ctx, tc, {{.DbTableLowerCamel}})
	})
	{{- else if eq .InterfaceType "修改"}}
	return daogext.Write(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		{{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return err
		}

		converter.{{.DbTableUpperCamel}}Converter.FillEntityWith{{.RequestModelName}}({{.DbTableLowerCamel}}, req)
		return dal.{{.DbTableUpperCamel}}ExtDao.Modify(ctx, tc, {{.DbTableLowerCamel}})
	})
	{{- else if eq .InterfaceType "删除"}}
	return daogext.Write(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		return dal.{{.DbTableUpperCamel}}ExtDao.DeleteById(ctx, tc, req.Id)
	})
	{{- else if eq .InterfaceType "分页"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		queryParam := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Param(req)
		pl, err := dal.{{.DbTableUpperCamel}}ExtDao.Page(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList(pl.List, converter.{{.DbTableUpperCamel}}Converter.Entity2{{.ResponseModelName}})
		return page.ListOf(pl.PageNo, pl.PageSize, pl.TotalCount, listModels), nil
	})
	{{- else if eq .InterfaceType "列表"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		queryParam := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Param(req)
		list, err := dal.{{.DbTableUpperCamel}}ExtDao.List(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList(list, converter.{{.DbTableUpperCamel}}Converter.Entity2{{.ResponseModelName}})
		return listModels, nil
	})
	{{- else if eq .InterfaceType "详情"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		{{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return nil, err
		}

		detailModel := converter.{{.DbTableUpperCamel}}Converter.Entity2{{.ResponseModelName}}({{.DbTableLowerCamel}})
		return detailModel, nil
	}){{end}}
}
{{end}}
`

func init() {
	ServiceTpl = ServiceTpl + ServiceAppendTpl
}
