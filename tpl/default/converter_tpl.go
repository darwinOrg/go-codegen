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

func (c *{{.LowerCamelName}}Converter) CreateReq2Entity(req *model.Create{{.GoTable}}Req) *dal.{{.GoTable}} {
	return &dal.{{.GoTable}}{
		{{range .CreateColumns}}{{.GoName}}: req.{{.GoName}},
		{{end}}
	}
}

func (c *{{.LowerCamelName}}Converter) FillEntityWithModifyReq({{.LowerCamelName}} *dal.{{.GoTable}}, req *model.Modify{{.GoTable}}Req) {
	{{range .ModifyColumns}}{{$.LowerCamelName}}.{{.GoName}} = req.{{.GoName}}
	{{end}}
}

func (c *{{.LowerCamelName}}Converter) Entity2ListVo({{.LowerCamelName}} *dal.{{.GoTable}}) *model.{{.GoTable}}ListVo {
	return &model.{{.GoTable}}ListVo{
		{{range .QueryColumns}}{{.GoName}}: {{$.LowerCamelName}}.{{.GoName}},
		{{end}}
		{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name: enum.{{$.GoTable}}{{.GoName}}Map[{{$.LowerCamelName}}.{{.GoName}}],{{end}}
		{{end}}
	}
}

func (c *{{.LowerCamelName}}Converter) Entity2DetailResp({{.LowerCamelName}} *dal.{{.GoTable}}) *model.{{.GoTable}}DetailResp {
	return &model.{{.GoTable}}DetailResp{
		{{range .QueryColumns}}{{.GoName}}: {{$.LowerCamelName}}.{{.GoName}},
		{{end}}
		{{range .QueryColumns}}{{if .HasEnum}}{{.GoName}}Name: enum.{{$.GoTable}}{{.GoName}}Map[{{$.LowerCamelName}}.{{.GoName}}],{{end}}
		{{end}}
	}
}

func (c *{{.LowerCamelName}}Converter) QueryReq2Param(req *model.Query{{.GoTable}}Req) *dal.Query{{.GoTable}}Param {
	return &dal.Query{{.GoTable}}Param {
		{{range .QueryColumns}}{{.GoName}}: req.{{.GoName}},
		{{end}}
	}
}

`
