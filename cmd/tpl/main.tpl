package main

import (
	_ "github.com/oaago/cloud/logx"
	"github.com/oaago/cloud/op"
	_ "github.com/oaago/cloud/preload"
	v2app "github.com/oaago/server/v2/app"
	"%package%/internal/consts"
	docs "%package%/docs"
	"%package%/internal/router"
)

func main() {
	application := v2app.Application{}
	app := application.Create()
	app.Router.SetTrustedProxies([]string{"localhost:9932"})
	app.HttpEngine.SetBaseUrl(app.Config.Server.BasePath)
	router.LoadRouterMapV2(app.HttpEngine)
	docs.SwaggerInfo.BasePath = app.Config.Server.BasePath
	app.HttpEngine.Options.HttpCode = consts.CODE
	event.EventInit(app.EventBus)
	app.Start()
}