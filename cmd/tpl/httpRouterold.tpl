package router
import (
	"github.com/oaago/server/oaa"
	{{if .HasMid}} middleware_http "{{.Module}}/internal/middleware/http"{{end}}
    {{range $index, $item := .MapHandlerMapImport}}
	{{$item.HttpDir}}{{$item.UpMethod}} "{{$item.Module}}/internal/api/{{$item.HttpDir}}/{{$item.Method}}" //{{$item}}{{end}}
)
type ConfigRouter oaa.ConfigRouter

func LoadRouterMap() oaa.MapHttpRoute {
	{{if eq .MiddlewareLen 0}}
	// 不需要中间件 Pid := middleware_http.NewPid()
	{{else}}
	Pid := middleware_http.NewPid()
	{{end}}
	m := oaa.MapHttpRoute{ {{range $index, $item := .MapHandlerMap}}
	"{{$item.Url}}": {
		{{range $index, $it := $item.Middleware}}Pid.{{$it}}{{end}}{{$item.HttpDir}}{{$item.UpMethod}}.{{$item.Upmet}}{{$item.Handler}}Handler,
	},{{end}}
	}
	return m
}