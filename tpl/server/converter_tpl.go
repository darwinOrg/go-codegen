package _server

var ConverterTpl = `package converter

import (
	"{{.PackagePrefix}}/dal"
	"{{.PackagePrefix}}/model"
	{{if .HasEnum}}"{{.PackagePrefix}}/enum"{{end}}
	"github.com/rolandhe/daog/ttypes"
)

var {{.DbTableUpperCamel}}Converter = &{{.DbTableLowerCamel}}Converter{}

type {{.DbTableLowerCamel}}Converter struct {
}
`

var ConverterAppendTpl = `
{{range .Requests}}
{{- if eq .InterfaceType "新建"}}
func (c *{{$.DbTableLowerCamel}}Converter) {{.UpperCamelName}}2Entity(req *model.{{.UpperCamelName}}) *dal.{{$.DbTableUpperCamel}} {
	if req == nil {
		return &dal.{{$.DbTableUpperCamel}}{}
	}

	{{$.DbTableLowerCamel}} := &dal.{{$.DbTableUpperCamel}}{
		{{range .Models}}{{.FieldName}}: {{if .NullableString}}*ttypes.FromString(req.{{.FieldName}}){{else}}req.{{.FieldName}}{{end}},
		{{end}}
	}

	return {{$.DbTableLowerCamel}}
}
{{else if eq .InterfaceType "修改"}}
func (c *{{$.DbTableLowerCamel}}Converter) FillEntityWith{{.UpperCamelName}}({{$.DbTableLowerCamel}} *dal.{{$.DbTableUpperCamel}}, req *model.{{.UpperCamelName}}) {
	if req == nil {
		return
	}

	{{range .Models}}{{if ne .FieldName "Id"}}{{$.DbTableLowerCamel}}.{{.FieldName}} = {{if .NullableString}}*ttypes.FromString(req.{{.FieldName}}){{else}}req.{{.FieldName}}{{end}}{{end}}
	{{end}}
}
{{else if .IsPageOrList}}
func (c *{{$.DbTableLowerCamel}}Converter) {{.UpperCamelName}}2Param(req *model.{{.UpperCamelName}}) *dal.Query{{$.DbTableUpperCamel}}Param {
	if req == nil {
		return &dal.Query{{$.DbTableUpperCamel}}Param{}
	}

	queryParam := &dal.Query{{$.DbTableUpperCamel}}Param {
		{{range .Models}}{{.FieldName}}: req.{{.FieldName}},
		{{end}}
		{{if .IsPage -}}PageParam: req.PageParam,{{end}}
	}

	return queryParam
}
{{end}}
{{end -}}
{{range .Responses}}
func (c *{{$.DbTableLowerCamel}}Converter) Entity2{{.UpperCamelName}}({{$.DbTableLowerCamel}} *dal.{{$.DbTableUpperCamel}}) *model.{{.UpperCamelName}} {
	if {{$.DbTableLowerCamel}} == nil {
		return &model.{{.UpperCamelName}}{}
	}

	return &model.{{.UpperCamelName}}{
		{{range .Models}}{{.FieldName}}: {{if ne .EnumFieldName ""}}enum.{{.EnumModel}}Map[{{$.DbTableLowerCamel}}.{{.EnumFieldName}}]{{else}}{{$.DbTableLowerCamel}}.{{.FieldName}}{{if .NullableString}}.StringNilAsEmpty(){{end}}{{end}},
		{{end}}
	}
}
{{end}}
`

func init() {
	ConverterTpl = ConverterTpl + ConverterAppendTpl
}
