package router
import (
	v2 "github.com/oaago/server/v2"
	{{if .HasMid}} middleware_http "{{.Module}}/internal/middleware/http"{{end}}
    {{range $index, $item := .MapHandlerMapImport}}
	{{$item.HttpDir}}{{$item.UpMethod}} "{{$item.Module}}/internal/api/{{$item.HttpDir}}/{{$item.Method}}" //{{$item}}{{end}}
)
func LoadRouterMapV2(http *v2.HttpEngine) {
	{{range $index, $item := .MapHandlerMap}}
	http.Router.{{$item.RequestType}}("{{$item.RequestUrl}}",{{range $index, $it := $item.Middleware}}Pid.{{$it}},{{end}} v2.NewHandler({{$item.HttpDir}}{{$item.UpMethod}}.{{$item.Upmet}}{{$item.Handler}}Handler)){{end}}
}