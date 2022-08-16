package {{.Package}}_{{.Method}}
import (
    "encoding/json"
    "github.com/go-playground/validator/v10"
    "github.com/oaago/cloud/logx"
    "github.com/oaago/server/oaa/translator"
    "github.com/oaago/server/v2"
	{{.Package}}{{.UpMethod}} "{{.Module}}/internal/service/{{.Package}}/{{.Method}}"
)

{{range $i, $method := .Upmet }}//{{$method}}{{$.UpPackage}}{{$.UpMethod}}Handler {{range $met, $msg := $.DecMessage }}{{if eq $met $method}}
//@Summary {{$.Dec}}
// @Schemes
// @Description {{$msg}}{{end}}{{end}}
// @Tags v1.0
// @Accept json
// @Produce json
{{range $.Param}}{{.}}
{{end}}// @Success 200 {object} {{$method}}{{$.UpPackage}}{{$.UpMethod}}Res
// @Router /{{$.Package}}/{{$.Method}} [{{$method}}]
func {{$method}}{{$.UpPackage}}{{$.UpMethod}}Handler(c *v2.Context) {
	// 实例化service
	{{$.Package}}Srv := {{$.Package}}{{$.UpMethod}}.NewService{{$.UpPackage}}()
	{{if eq $method "Post" "Put"}}
	if err := c.ShouldBindJSON(&{{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Req); err != nil {
	{{else}}
	if err := c.ShouldBind(&{{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Req); err != nil {
	{{end}}
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.Return(200, nil, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.Return(200, nil, errs.Translate(translator.Trans))
		return
	}
    // 打印相关信息
    req, _ := json.Marshal({{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Req)
    res, _ := json.Marshal({{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Res)
    logx.Logger.Info(`{{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Req: `+ string(req))
    logx.Logger.Info(`{{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Res: `+ string(res))
	{{if eq $method "Get"}}
	var err = {{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Service()
	{{else if eq $method "Delete"}}
	var err = {{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Service()
	{{else}}
	_, err := {{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Service(){{end}}
	if err != nil {
		logx.Logger.Info("{{$.UpPackage}}{{$.UpMethod}}数据请求异常", {{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Res, err.Error())
		c.Return(10006)
		return
	}
	logx.Logger.Info("{{$.UpPackage}}{{$.UpMethod}}数据请求正常", {{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Res)
	c.Return(200, {{$.Package}}Srv.{{$method}}{{$.UpPackage}}{{$.UpMethod}}Res)
	return
}

{{end}}