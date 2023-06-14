package router
import (
	http "github.com/oaago/server/v2/http/bootstrap"
	"github.com/oaago/server/v2/types"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	middlewarehttp "oaatpl/internal/middleware/http"
	{{if .HasMid}}middlewarehttp "{{.Module}}/internal/middleware/http"{{end}}
    {{range $index, $item := .MapHandlerMapImport}}
	{{$item.HttpDir}}{{$item.UpMethod}} "{{$item.Module}}/internal/api/{{$item.HttpDir}}/{{$item.Method}}" //{{$item}}{{end}}
)
func LoadRouterMapV2(h *types.HttpEngine) {
    h.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	{{if .HasMid}}
	 // 示例化局部中间件
	 Pid := middlewarehttp.NewPid()
	{{end}}
	v := h.Router.Group(h.Options.BaseUrl)
	{
	{{range $index, $item := .MapHandlerMap}}{{if $item.RN}}

	    // {{$item.HttpDir}}{{$item.UpMethod}} {{$item.RequestUrl}}{{end}}
	   v.{{$item.RequestType}}("{{$item.RequestUrl}}", {{range $index, $it := $item.Middleware}}http.NewHandler(Pid.{{$it}}),{{end}} http.NewHandler({{$item.HttpDir}}{{$item.UpMethod}}.{{$item.Handler}})){{end}}
    }
}