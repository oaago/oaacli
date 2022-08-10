package main

import (
	_ "github.com/oaago/cloud/logx"
	"github.com/oaago/cloud/op"
	_ "github.com/oaago/cloud/preload"
	"github.com/oaago/server/oaa"
	"%package%/internal/consts"
	docs "%package%/docs"
	"%package%/internal/router"
)

func main() {
	op.ConfigData.CodeMap = consts.CODE
	route := &router.ConfigRouter{
		MapHttpRoute: router.LoadRouterMap,
	}
	rr := oaa.NewRouter((*oaa.ConfigRouter)(route))
	//route.RpcServer = route.RegisterRpcGenRouter()
	docs.SwaggerInfo.BasePath = op.ConfigData.Server.BasePath
	oaa.Start(rr)
}