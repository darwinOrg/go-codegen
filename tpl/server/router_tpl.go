package _server

var RouterTpl = `package router

import (
	"{{.PackagePrefix}}/handler"
	{{if .HasModel}}"{{.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	"github.com/darwinOrg/go-common/result"
	"github.com/darwinOrg/go-web/wrapper"
	"github.com/gin-gonic/gin"
)

func Bind{{.GroupUpperCamel}}Router(rg *gin.RouterGroup) {
	g := rg.Group("{{.RoutePrefix}}")
	{{range .Models}}
	wrapper.{{.MethodType}}(&wrapper.RequestHolder[{{.RequestModelNameExp}}, *result.Result[{{.ResponseModelNameExp}}]]{
		Remark: "{{.Remark}}", RouterGroup: g, RelativePath: "{{.RelativePath}}",
		{{if .NonLogin}}NonLogin: true, {{- end -}}
		{{if .NotLogSQL}}NotLogSQL: true, {{- end -}}
		{{if ne .LogLevelExp ""}}LogLevel: {{.LogLevelExp}}, {{- end}}
		BizHandler: handler.{{$.GroupUpperCamel}}Handler.{{.MethodNameExp}},
	})
	{{end}}
}
`
