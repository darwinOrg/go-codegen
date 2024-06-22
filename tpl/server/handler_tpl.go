package _server

var HandlerTpl = `package handler

import (
	{{if .HasModel}}"{{.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	"{{.PackagePrefix}}/service"
	dgctx "github.com/darwinOrg/go-common/context"
	"github.com/darwinOrg/go-common/result"
	"github.com/gin-gonic/gin"
)
{{range $index, $inter := .Interfaces}}
var {{$inter.GroupUpperCamel}}Handler = &{{$inter.GroupLowerCamel}}Handler{}
type {{$inter.GroupLowerCamel}}Handler struct {
}
{{range .Models}}
func (h *{{$inter.GroupLowerCamel}}Handler) {{.MethodNameExp}}(_ *gin.Context, ctx *dgctx.DgContext, req *{{.RequestModelNameExp}}) *result.Result[{{.ResponseModelNameExp}}] {
	{{- if ne .ResponseModelName ""}}
	resp, err := service.{{$inter.Group}}Service.{{.MethodNameExp}}(ctx, req)
	if err != nil {
		return result.SimpleFail[{{.ResponseModelNameExp}}](err.Error())
	}
	return result.Success[{{.ResponseModelNameExp}}](resp)
	{{- else}}
	err := service.{{$inter.Group}}Service.{{.MethodNameExp}}(ctx, req)
	if err != nil {
		return result.SimpleFailByError(err)
	}
	return result.SimpleSuccess()
	{{end -}}
}
{{end}}
{{end}}
`
