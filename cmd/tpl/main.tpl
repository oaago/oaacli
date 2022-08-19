package main

import (
	_ "github.com/oaago/cloud/logx"
	"github.com/oaago/cloud/op"
	_ "github.com/oaago/cloud/preload"
	"github.com/oaago/server/v2"
	"%package%/internal/consts"
	docs "%package%/docs"
	"%package%/internal/router"
)

func main() {
	op.ConfigData.CodeMap = consts.CODE
	ops := v2.HttpConfig{
		Port: op.ConfigData.Server.Port,
	}
	http := v2.NewRouter(ops)
	http.SetBaseUrl(op.ConfigData.Server.BasePath)
	router.LoadRouterMapV2(http)
	docs.SwaggerInfo.BasePath = op.ConfigData.Server.BasePath
	http.Start()
}