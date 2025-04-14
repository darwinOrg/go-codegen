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
var {{.GroupUpperCamel}}Handler = &{{.GroupLowerCamel}}Handler{}
type {{.GroupLowerCamel}}Handler struct {
}
`

var HandlerAppendTpl = `
{{range .Models}}
func (h *{{$.GroupLowerCamel}}Handler) {{.MethodNameExp}}(_ *gin.Context, ctx *dgctx.DgContext, req *{{.RequestModelNameExp}}) *result.Result[{{.ResponseModelNameExp}}] {
	{{- if ne .ResponseModelName ""}}
	resp, err := service.{{$.GroupUpperCamel}}Service.{{.MethodNameExp}}(ctx, req)
	if err != nil {
		return result.FailByError[{{.ResponseModelNameExp}}](err)
	}
	return result.Success[{{.ResponseModelNameExp}}](resp)
	{{- else}}
	err := service.{{$.GroupUpperCamel}}Service.{{.MethodNameExp}}(ctx, req)
	if err != nil {
		return result.SimpleFailByError(err)
	}
	return result.SimpleSuccess(){{end}}
}
{{end}}
`

func init() {
	HandlerTpl = HandlerTpl + HandlerAppendTpl
}
