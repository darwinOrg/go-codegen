package _client

var ClientTpl = `package client

import (
	{{if .HasModel}}"{{.PackagePrefix}}/model"{{end}}
	{{if .HasId}}cm "github.com/darwinOrg/go-common/model"{{end}}
	{{if .HasPage}}"github.com/darwinOrg/go-common/page"{{end}}
	dgctx "github.com/darwinOrg/go-common/context"
	"github.com/darwinOrg/go-common/result"
	dghttp "github.com/darwinOrg/go-httpclient"
)
{{range $index, $inter := .Interfaces}}
type {{$inter.GroupLowerCamel}}Client struct {
	Host string
}
{{range .Models}}
func (c *{{$inter.GroupLowerCamel}}Client) {{.MethodNameExp}}(ctx *dgctx.DgContext, req *{{.RequestModelNameExp}}) {{if ne .ResponseModelName ""}}({{.ResponseModelNameExp}}, error){{else}}error{{end}} {
	urlPath := "http://" + c.Host + "{{$inter.RoutePrefix}}/{{.RelativePath}}"
	rt, err := dghttp.DoPostJsonToResult[{{.ResponseModelNameExp}}](ctx, urlPath, req, nil)
	if err != nil {
		{{if ne .ResponseModelName ""}}rt = result.FailByError[{{.ResponseModelNameExp}}](err){{else}}return err{{end}}
	}
	{{if ne .ResponseModelName ""}}return result.ExtractData(rt){{else}}return result.ToError(rt){{end}}
}
{{end}}
{{end}}
`
