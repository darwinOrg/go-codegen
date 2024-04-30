package tpl

var ConverterExtTpl = `package converter

import (
	"{{.ProjectPath}}/dal"
	"{{.ProjectPath}}/model"
)

var {{.GoTable}}Converter = &{{.LowerCamelName}}Converter{}

type {{.LowerCamelName}}Converter struct {
}

func (c *{{.LowerCamelName}}Converter) CreateModel2Entity(req *model.Create{{.GoTable}}Req) *dal.{{.GoTable}} {
	{{.LowerCamelName}} := &dal.{{.GoTable}}{
		{{range .CreateColumns}}{{.GoName}}: req.{{.GoName}},
		{{end}}
	}

	return {{.LowerCamelName}}
}

func (c *{{.LowerCamelName}}Converter) ModifyModel2Entity(req *model.Modify{{.GoTable}}Req) *dal.{{.GoTable}} {
	{{.LowerCamelName}} := &dal.{{.GoTable}}{
		{{range .ModifyColumns}}{{.GoName}}: req.{{.GoName}},
		{{end}}
	}

	return {{.LowerCamelName}}
}

func (c *{{.LowerCamelName}}Converter) FillEntityWithModifyModel(entity *dal.{{.GoTable}}, req *model.Modify{{.GoTable}}Req) {
	{{range .ModifyColumns}}entity.{{.GoName}} = req.{{.GoName}}
	{{end}}
}

func (c *{{.LowerCamelName}}Converter) Entity2ListModel(entity *dal.{{.GoTable}}) *model.{{.GoTable}}ListResp {
	listVo := &model.{{.GoTable}}ListResp{
		{{range .QueryColumns}}{{.GoName}}: entity.{{.GoName}},
		{{end}}
	}

	return listVo
}

func (c *{{.LowerCamelName}}Converter) Entity2DetailModel(entity *dal.{{.GoTable}}) *model.{{.GoTable}}DetailResp {
	detailVo := &model.{{.GoTable}}DetailResp{
		{{range .QueryColumns}}{{.GoName}}: entity.{{.GoName}},
		{{end}}
	}

	return detailVo
}

`
