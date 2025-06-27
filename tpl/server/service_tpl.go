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
func (s *{{$.GroupLowerCamel}}Service) {{.MethodNameExp}}(ctx *dgctx.DgContext, req *{{.RequestModelNameExp}}) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
	{{- if eq .DbModelName ""}}
	// TODO
	return {{if ne .ResponseModelName ""}}nil, {{end}}nil
	{{- else if eq .InterfaceType "新建"}}
	return daogext.Write{{if ne .ResponseModelName ""}}WithResult{{end}}(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		{{.DbTableLowerCamel}} := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Entity(req)
		return dal.{{.DbTableUpperCamel}}ExtDao.Create(ctx, tc, {{.DbTableLowerCamel}})
	})
	{{- else if eq .InterfaceType "修改"}}
	return daogext.Write{{if ne .ResponseModelName ""}}WithResult{{end}}(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		{{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return err
		}

		converter.{{.DbTableUpperCamel}}Converter.FillEntityWith{{.RequestModelName}}({{.DbTableLowerCamel}}, req)
		return dal.{{.DbTableUpperCamel}}ExtDao.Modify(ctx, tc, {{.DbTableLowerCamel}})
	})
	{{- else if eq .InterfaceType "保存"}}
	return daogext.Write{{if ne .ResponseModelName ""}}WithResult{{end}}(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		if req.Id > 0 {
			{{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
			if err != nil {
				return err
			}
	
			converter.{{.DbTableUpperCamel}}Converter.FillEntityWith{{.RequestModelName}}({{.DbTableLowerCamel}}, req)
			return dal.{{.DbTableUpperCamel}}ExtDao.Modify(ctx, tc, {{.DbTableLowerCamel}})
		} else {
			{{.DbTableLowerCamel}} := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Entity(req)
			return dal.{{.DbTableUpperCamel}}ExtDao.Create(ctx, tc, {{.DbTableLowerCamel}})
		}
	})
	{{- else if eq .InterfaceType "删除"}}
	return daogext.Write{{if ne .ResponseModelName ""}}WithResult{{end}}(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		return dal.{{.DbTableUpperCamel}}ExtDao.DeleteById(ctx, tc, req.Id)
	})
	{{- else if eq .InterfaceType "分页"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		queryParam := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Param(req)
		pl, err := dal.{{.DbTableUpperCamel}}ExtDao.Page(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listVos := dgcoll.MapToList(pl.List, converter.{{.DbTableUpperCamel}}Converter.Entity2{{.ResponseModelName}})
		return page.ListOf(pl.PageNo, pl.PageSize, pl.TotalCount, listVos), nil
	})
	{{- else if eq .InterfaceType "列表"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		queryParam := converter.{{.DbTableUpperCamel}}Converter.{{.RequestModelName}}2Param(req)
		list, err := dal.{{.DbTableUpperCamel}}ExtDao.List(ctx, tc, queryParam)
		if err != nil {
			return nil, err
		}

		listVos := dgcoll.MapToList(list, converter.{{.DbTableUpperCamel}}Converter.Entity2{{.ResponseModelName}})
		return listVos, nil
	})
	{{- else if eq .InterfaceType "详情"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		{{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return nil, err
		}

		detailResp := converter.{{.DbTableUpperCamel}}Converter.Entity2{{.ResponseModelName}}({{.DbTableLowerCamel}})
		return detailResp, nil
	})
	{{- else if eq .InterfaceType "只读"}}
	return daogext.Readonly{{if ne .ResponseModelName ""}}WithResult{{end}}(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		// {{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		_, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return {{if ne .ResponseModelName ""}}nil, {{end}}err
		}

		// TODO
		return {{if ne .ResponseModelName ""}}nil, {{end}}err
	})
	{{- else if eq .InterfaceType "其它"}}
	return daogext.Write{{if ne .ResponseModelName ""}}WithResult{{end}}(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		// {{.DbTableLowerCamel}}, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		_, err := dal.{{.DbTableUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return {{if ne .ResponseModelName ""}}nil, {{end}}err
		}

		// TODO
		return {{if ne .ResponseModelName ""}}nil, {{end}}err
	})
	{{end}}
}
{{end}}
`

func init() {
	ServiceTpl = ServiceTpl + ServiceAppendTpl
}
