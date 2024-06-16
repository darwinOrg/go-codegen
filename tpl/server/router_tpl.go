package _server

var RouterExtTpl = `package router

import (
	"{{.Export.PackagePrefix}}/handler"
	{{if .HasModel}}"{{.Export.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	"github.com/darwinOrg/go-common/result"
	"github.com/darwinOrg/go-web/wrapper"
	"github.com/gin-gonic/gin"
)

{{range $index, $inter := .Interfaces}}
func Bind{{$inter.Group}}Router(rg *gin.RouterGroup) {
	g := rg.Group("{{.RoutePrefix}}")
	{{range .Models}}
	wrapper.Post(&wrapper.RequestHolder[{{if eq .RequestModelName "Id"}}cm.IdReq{{else if ne .RequestModelName ""}}model.{{.RequestModelName}}{{else}}result.Void{{end}}, *result.Result[{{if ne .ResponseModelName ""}}{{if .ResponseModelHasPointer}}*{{end}}model.{{.ResponseModelName}}{{else}}*result.Void{{end}}]]{
		Remark:       "{{.Remark}}",
		RouterGroup:  g,
		RelativePath: "{{.RelativePath}}",
		NonLogin: 	  {{.NonLogin}},
		NotLogSQL:    {{.NotLogSQL}},
		{{if ne .AllowProductsExp ""}}AllowProducts: {{.AllowProductsExp}},{{- end -}}
		{{if ne .LogLevelExp ""}}LogLevel: {{.LogLevelExp}},{{- end}}
		BizHandler:   handler.{{$inter.Group}}Handler.{{.MethodNameExp}},
	})
	{{end}}
}
{{end}}
`
