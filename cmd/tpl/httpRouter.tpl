package router
import (
	v2 "github.com/oaago/server/v2"
	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	{{if .HasMid}}middlewarehttp "{{.Module}}/internal/middleware/http"{{end}}
    {{range $index, $item := .MapHandlerMapImport}}
	{{$item.HttpDir}}{{$item.UpMethod}} "{{$item.Module}}/internal/api/{{$item.HttpDir}}/{{$item.Method}}" //{{$item}}{{end}}
)
func LoadRouterMapV2(http *v2.HttpEngine) {
    http.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	{{if .HasMid}}
	 // 示例化局部中间件
	 Pid := middlewarehttp.NewPid()
	{{end}}
	v := http.Router.Group(http.Options.BaseUrl)
	{
	{{range $index, $item := .MapHandlerMap}}{{if $item.RN}}

	    // {{$item.HttpDir}}{{$item.UpMethod}} {{$item.RequestUrl}}{{end}}
	   v.{{$item.RequestType}}("{{$item.RequestUrl}}", {{range $index, $it := $item.Middleware}}v2.NewHandler(Pid.{{$it}}),{{end}} v2.NewHandler({{$item.HttpDir}}{{$item.UpMethod}}.{{$item.Handler}})){{end}}
    }
}