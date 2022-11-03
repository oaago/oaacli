package main

import (
	_ "github.com/oaago/cloud/logx"
	"github.com/oaago/cloud/op"
	_ "github.com/oaago/cloud/preload"
	"github.com/oaago/server/v2/app"
	h "github.com/oaago/server/v2/http"
	"%package%/internal/consts"
	docs "%package%/docs"
	"%package%/internal/router"
)

func main() {
	op.ConfigData.CodeMap = consts.CODE
	ops := app.HttpConfig{
		Port: op.ConfigData.Server.Port,
	}
	http := h.NewRouter(ops)
	http.SetBaseUrl(op.ConfigData.Server.BasePath)
	router.LoadRouterMapV2(http)
	docs.SwaggerInfo.BasePath = op.ConfigData.Server.BasePath
	http.Start()
}