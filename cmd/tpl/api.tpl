package {{.Package}}_{{.Method}}
import (
    "encoding/json"
    "github.com/oaago/cloud/logx"
    "github.com/oaago/server/v2"
    "github.com/oaago/server/v2/translator"
	{{.Package}}{{.UpMethod}} "{{.Module}}/internal/service/{{.Package}}/{{.Method}}"
)

{{range $Handler := .HandlerContent }}
// {{$Handler.HandlerName}}
// @Summary {{$Handler.Dec}}
// @Schemes
// @Description {{$Handler.Dec}}
// @Tags {{$Handler.Comment}}-{{$.UpPackage}}{{$.UpMethod}}
// @Accept json
// @Produce json
{{- range $p := $Handler.Param}}
{{$p}}
{{- end}}
// @Success 200 {object} {{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Res
// @Router /{{$.Package}}/{{$.Method}} [{{$Handler.Method}}]
func {{$Handler.HandlerName}}(c *v2.Context) {
	// 实例化service
	{{$.Package}}Srv := {{$.Package}}{{$.UpMethod}}.NewService{{$.UpPackage}}()
    // 绑定参数
	{{if eq $Handler.Method "Get"}}if err := c.ShouldBindQuery(&{{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Req); err != nil {
	{{else}} if err := c.ShouldBindJSON(&{{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Req); err != nil {
	{{end}}c.Return(200, nil, err.Error())
        return
	}
    // 再先验证
    errs := translator.InitTrans({{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Req)
    if errs != nil {
        c.Return(200, nil, errs)
        return
    }
    // 打印相关信息
    req, _ := json.Marshal({{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Req)
    res, _ := json.Marshal({{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Res)
    logx.Logger.Info(`{{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Req: `+ string(req))
    logx.Logger.Info(`{{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Res: `+ string(res))
    {{if eq $Handler.Method "Get"}}
	var err = {{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Service(){{else if eq $Handler.Method "Delete"}}
	var err = {{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Service(){{else}}
	_, err := {{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Service(){{end}}
	if err != nil {
		logx.Logger.Info("{{$.UpPackage}}{{$.UpMethod}}数据请求异常", {{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Res, err.Error())
		c.Return(10006)
		return
	}
	logx.Logger.Info("{{$.UpPackage}}{{$.UpMethod}}数据请求正常", {{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Res)
	c.Return(200, {{$.Package}}Srv.{{$Handler.Method}}{{$.UpPackage}}{{$.UpMethod}}Res)
	return
}

{{end}}