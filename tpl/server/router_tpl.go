package _server

var RouterExtTpl = `package router

import (
	"{{.PackagePrefix}}/handler"
	{{if .HasModel}}"{{.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	"github.com/darwinOrg/go-common/result"
	"github.com/darwinOrg/go-web/wrapper"
	"github.com/gin-gonic/gin"
)
{{range $index, $inter := .Interfaces}}
func Bind{{$inter.GroupUpperCamel}}Router(rg *gin.RouterGroup) {
	g := rg.Group("{{.RoutePrefix}}")
	{{range .Models}}
	wrapper.Post(&wrapper.RequestHolder[{{.RequestModelNameExp}}, *result.Result[{{.ResponseModelNameExp}}]]{
		Remark:       "{{.Remark}}", RouterGroup:  g, RelativePath: "{{.RelativePath}}",
		{{if .NonLogin}}NonLogin: true,{{- end -}}
		{{if .NotLogSQL}}NotLogSQL: true,{{- end -}}
		{{if ne .AllowProductsExp ""}}AllowProducts: {{.AllowProductsExp}},{{- end -}}
		{{if ne .LogLevelExp ""}}LogLevel: {{.LogLevelExp}},{{- end}}
		BizHandler:   handler.{{$inter.GroupUpperCamel}}Handler.{{.MethodNameExp}},
	})
	{{end}}
}
{{end}}
`
