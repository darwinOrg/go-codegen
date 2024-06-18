package _server

var ServiceExtTpl = `package service

import (
	{{if .HasModel}}"{{.Export.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	{{if .HasQuery}}dgcoll "github.com/darwinOrg/go-common/collection"{{end}}
	{{if .HasModel}}"{{.Export.PackagePrefix}}/converter"{{end}}
	"{{.Export.PackagePrefix}}/dal"
	daogext "github.com/darwinOrg/daog-ext"
	dgctx "github.com/darwinOrg/go-common/context"
	"github.com/rolandhe/daog"
)
{{range $index, $inter := .Interfaces}}
var {{$inter.GroupUpperCamel}}Service = &{{$inter.GroupLowerCamel}}Service{}
type {{$inter.GroupLowerCamel}}Service struct {
}
{{range .Models}}
func (s *{{$inter.GroupLowerCamel}}Service) {{.MethodNameExp}}(ctx *dgctx.DgContext, req *{{.RequestModelNameExp}}) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
	{{- if eq .InterfaceType "新建"}}
	entity := converter.{{$inter.GroupUpperCamel}}Converter.{{.RequestModelName}}2Entity(req)
	
	return daogext.WriteWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		return dal.{{$inter.GroupUpperCamel}}ExtDao.Create(ctx, tc, entity)
	})
	{{- else if eq .InterfaceType "修改"}}
	return daogext.Write(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		entity, err := dal.{{$inter.GroupUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return err
		}

		converter.{{$inter.GroupUpperCamel}}Converter.FillEntityWith{{.RequestModelName}}(entity, req)
		return dal.{{$inter.GroupUpperCamel}}ExtDao.Modify(ctx, tc, entity)
	})
	{{- else if eq .InterfaceType "删除"}}
	return daogext.Write(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		return dal.{{$inter.GroupUpperCamel}}ExtDao.DeleteById(ctx, tc, req.Id)
	})
	{{- else if eq .InterfaceType "分页"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		pl, err := dal.{{$inter.GroupUpperCamel}}ExtDao.Page(ctx, tc, req)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList(pl.List, converter.{{$inter.GroupUpperCamel}}Converter.Entity2{{.ResponseModelName}})
		return page.ListOf[model.{{.ResponseModelName}}]](pl.PageNo, pl.PageSize, pl.TotalCount, listModels), nil
	})
	{{- else if eq .InterfaceType "列表"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		list, err := dal.{{$inter.GroupUpperCamel}}ExtDao.List(ctx, tc, req)
		if err != nil {
			return nil, err
		}

		listModels := dgcoll.MapToList(list, converter.{{$inter.GroupUpperCamel}}Converter.Entity2{{.ResponseModelName}})
		return listModels, nil
	})
	{{- else if eq .InterfaceType "详情"}}
	return daogext.ReadonlyWithResult(ctx, func(tc *daog.TransContext) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
		entity, err := dal.{{$inter.GroupUpperCamel}}ExtDao.MustGetById(ctx, tc, req.Id)
		if err != nil {
			return nil, err
		}

		detailModel := converter.{{$inter.GroupUpperCamel}}Converter.Entity2{{.ResponseModelName}}(entity)
		return detailModel, nil
	})
	{{end}}
}
{{end}}
{{end}}
`
