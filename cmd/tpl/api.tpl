package {{.Package}}
import (
    "encoding/json"
    "github.com/oaago/cloud/logx"
    "github.com/oaago/server/v2/http"
    "github.com/oaago/server/v2/http/translator"
	{{.Package}} "{{.Module}}/internal/service/{{.ServicePath}}"
)
// {{.HandlerName}}Handler
// @Summary {{.Dec}}
// @Schemes
// @Description {{.Dec}}
// @Tags {{.Comment}}-{{.UpPackage}}{{.UpMethod}}
// @Accept json
// @Produce json
{{- range $p := .Param}}
{{$p}}
{{- end}}
// @Success 200 {object} {{.HandlerName}}Res
// @Router /{{.ServicePath}}/{{.HandlerName}} [{{.Method}}]
func {{.HandlerName}}Handler(c *http.Context) {
	// 实例化service
	{{.ServiceName}}Srv := {{.Package}}.NewService{{.ServiceName}}()
    // 绑定参数
	{{if eq .Method "Get"}}if err := c.ShouldBindQuery(&{{.ServiceName}}Srv.{{.HandlerName}}Req); err != nil {
	{{else}} if err := c.ShouldBindJSON(&{{.ServiceName}}Srv.{{.HandlerName}}Req); err != nil {
	{{end}}c.Return(200, nil, err.Error())
        return
	}
    // 再先验证
    errs := translator.InitTrans({{.ServiceName}}Srv.{{.HandlerName}}Req)
    if errs != nil {
        c.Return(200, nil, errs)
        return
    }
    // 打印相关信息
    req, _ := json.Marshal({{.ServiceName}}Srv.{{.HandlerName}}Req)
    res, _ := json.Marshal({{.ServiceName}}Srv.{{.HandlerName}}Res)
    logx.Logger.Info(`{{.ServiceName}}Srv.{{.HandlerName}}Req: `+ string(req))
    logx.Logger.Info(`{{.ServiceName}}Srv.{{.HandlerName}}Res: `+ string(res))
    var ResErr = {{.ServiceName}}Srv.{{.HandlerName}}Service()
	if ResErr != nil {
		logx.Logger.Info("{{.UpPackage}}{{.UpMethod}}数据请求异常", {{.ServiceName}}Srv.{{.HandlerName}}Res, ResErr.Error())
		c.Return(10006, nil, ResErr.Error())
		return
	}
	logx.Logger.Info("{{.UpPackage}}{{.UpMethod}}数据请求正常", {{.ServiceName}}Srv.{{.HandlerName}}Res)
	c.Return(200, {{.ServiceName}}Srv.{{.HandlerName}}Res)
	return
}