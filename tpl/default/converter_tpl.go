package _default

var ConverterTpl = `package converter

import (
	"{{.PackagePrefix}}/dal"
	"{{.PackagePrefix}}/model"
	{{if .HasEnum}}"{{.PackagePrefix}}/enum"{{end}}
)

var {{.GoTable}}Converter = &{{.LowerCamelName}}Converter{}

type {{.LowerCamelName}}Converter struct {
}

func (c *{{.LowerCamelName}}Converter) CreateModel2Entity(req *model.Create{{.GoTable}}Req) *dal.{{.GoTable}} {
	if req == nil {
		return nil
	}

	{{.LowerCamelName}} := &dal.{{.GoTable}}{
		{{range .CreateColumns}}{{.GoName}}: req.{{.GoName}},
		{{end}}
	}

	return {{.LowerCamelName}}
}

func (c *{{.LowerCamelName}}Converter) FillEntityWithModifyModel({{.LowerCamelName}} *dal.{{.GoTable}}, req *model.Modify{{.GoTable}}Req) {
	if {{.LowerCamelName}} == nil {
		return
	}

	{{range .ModifyColumns}}{{$.LowerCamelName}}.{{.GoName}} = req.{{.GoName}}
	{{end}}
}

func (c *{{.LowerCamelName}}Converter) Entity2ListModel({{.LowerCamelName}} *dal.{{.GoTable}}) *model.{{.GoTable}}ListResp {
	if {{.LowerCamelName}} == nil {
		return nil
	}

	listVo := &model.{{.GoTable}}ListResp{
		{{range .QueryColumns}}{{.GoName}}: {{$.LowerCamelName}}.{{.GoName}},
		{{end}}
		{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name: enum.{{$.GoTable}}{{.GoName}}Map[{{$.LowerCamelName}}.{{.GoName}}],{{end}}
		{{end}}
	}

	return listVo
}

func (c *{{.LowerCamelName}}Converter) Entity2DetailModel({{.LowerCamelName}} *dal.{{.GoTable}}) *model.{{.GoTable}}DetailResp {
	if {{.LowerCamelName}} == nil {
		return nil
	}

	detailVo := &model.{{.GoTable}}DetailResp{
		{{range .QueryColumns}}{{.GoName}}: {{$.LowerCamelName}}.{{.GoName}},
		{{end}}
		{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name: enum.{{$.GoTable}}{{.GoName}}Map[{{$.LowerCamelName}}.{{.GoName}}],{{end}}
		{{end}}
	}

	return detailVo
}

func (c *{{.LowerCamelName}}Converter) QueryModel2Param(req *model.Query{{.GoTable}}Req) *dal.Query{{.GoTable}}Param {
	return &dal.Query{{.GoTable}}Param {
		{{range .QueryColumns}}{{.GoName}}: req.{{.GoName}},
		{{end}}
	}
}

`
