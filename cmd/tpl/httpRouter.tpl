package router
import (
	v2 "github.com/oaago/server/v2"
	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	{{if .HasMid}} middleware_http "{{.Module}}/internal/middleware/http"{{end}}
    {{range $index, $item := .MapHandlerMapImport}}
	{{$item.HttpDir}}{{$item.UpMethod}} "{{$item.Module}}/internal/api/{{$item.HttpDir}}/{{$item.Method}}" //{{$item}}{{end}}
)
func LoadRouterMapV2(http *v2.HttpEngine) {
    http.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	{{range $index, $item := .MapHandlerMap}}
	http.Router.{{$item.RequestType}}("{{$item.RequestUrl}}",{{range $index, $it := $item.Middleware}}Pid.{{$it}},{{end}} v2.NewHandler({{$item.HttpDir}}{{$item.UpMethod}}.{{$item.Upmet}}{{$item.Handler}}Handler)){{end}}
}