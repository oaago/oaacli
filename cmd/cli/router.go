package cli

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	tpl "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
	"github.com/tidwall/gjson"
)

type MapHttpHandler struct {
	Url             string
	Handler         string
	Module          string
	Method          string
	UpMethod        string
	Package         string
	UpPackage       string
	HttpDir         string
	Middleware      []string
	Upmet           string
	HandlerMapOfOne map[string]bool
}

type HttpMap struct {
	MapHandlerMap       []MapHttpHandler
	MiddlewareLen       int
	Module              string
	MapHandlerMapImport []MapHttpHandler
}

func genRouter(module, handler, pack, url string) {
	MapHandlerMap := make([]MapHttpHandler, 0)
	MapHandlerMapImport := make([]MapHttpHandler, 0)
	HandlerMapOfOne := map[string]bool{}
	var httpMap HttpMap
	data, _ := os.ReadFile("./oaa.json")
	httpData := gjson.Get(string(data), "http").Array()
	fmt.Println(httpData, "路由信息")
	for _, datum := range httpData {
		lis := strings.Split(datum.String(), "@/")
		if len(lis) != 2 {
			panic(datum.Str + "不符合规则")
		}
		lim := strings.Split(lis[1], "|")
		dir := lim[0]
		hand := strings.Split(dir, "/")
		Middleware := make([]string, 0)
		if len(lim) > 1 {
			if strings.Contains(lim[1], ",") {
				midList := strings.Split(lim[1], ",")
				for _, mid := range midList {
					if len(mid) > 0 {
						Middleware = append(Middleware, mid+",")
					}
				}
			} else {
				Middleware = append(Middleware, lim[1]+",")
			}
		}
		methodMap := strings.Split(lis[0], ",")
		for _, s := range methodMap {
			// 判断有多少个handler 文件
			if HandlerMapOfOne[dir] != true {
				HandlerMapOfOne[dir] = true
				MapHandlerMapImport = append(MapHandlerMapImport, MapHttpHandler{
					Url:             s + "@/" + dir,
					Module:          module,
					Middleware:      Middleware,
					Handler:         utils.Ucfirst(hand[0]) + utils.Ucfirst(hand[1]),
					HttpDir:         utils.Camel2Case(hand[0]),
					Method:          hand[1],
					UpMethod:        utils.Ucfirst(hand[1]),
					Package:         hand[1],
					UpPackage:       utils.Camel2Case(pack),
					Upmet:           utils.Ucfirst(s),
					HandlerMapOfOne: HandlerMapOfOne,
				})
			}
			MapHandlerMap = append(MapHandlerMap, MapHttpHandler{
				Url:        s + "@/" + dir,
				Module:     module,
				Middleware: Middleware,
				Handler:    utils.Ucfirst(hand[0]) + utils.Ucfirst(hand[1]),
				HttpDir:    utils.Camel2Case(hand[0]),
				Method:     hand[1],
				UpMethod:   utils.Ucfirst(hand[1]),
				Package:    hand[1],
				UpPackage:  utils.Camel2Case(pack),
				Upmet:      utils.Ucfirst(s),
			})
		}
		if len(Middleware) != 0 {
			httpMap.MiddlewareLen = len(Middleware)
		}
	}
	tmpl, e := template.New("gen-http-router").Parse(strings.TrimSpace(tpl.HttpRouterTpl))
	if e != nil {
		panic(e.Error())
	}
	httpMap.MapHandlerMap = MapHandlerMap
	httpMap.Module = module
	httpMap.MapHandlerMapImport = MapHandlerMapImport
	httpFile, _ := os.Create(HttpRouterFile)
	if err := tmpl.Execute(httpFile, httpMap); err != nil {
		panic(err)
	}
	fmt.Println("http 路由文件生成成功")
}
func genRpcRouter(module, handler, pack, url string) {
	MapHandlerMap := make([]MapHttpHandler, 0)
	var httpMap HttpMap
	data, _ := os.ReadFile(configFile)
	RpcData := gjson.Get(string(data), "rpc").Array()
	for _, datum := range RpcData {
		lis := strings.Split(datum.String(), "&/")
		lim := strings.Split(lis[1], "|")
		fmt.Println(lim, lis)
		dir := lim[0]
		hand := strings.Split(dir, "/")
		MapHandlerMap = append(MapHandlerMap, MapHttpHandler{
			Url:       datum.String(),
			Module:    module,
			Handler:   utils.Ucfirst(hand[0]) + utils.Ucfirst(hand[1]),
			Method:    hand[1],
			UpMethod:  utils.Ucfirst(hand[1]),
			Package:   hand[0],
			UpPackage: utils.Ucfirst(hand[0]),
		})
	}
	tmpl, e := template.New("gen-rpc-router").Parse(strings.TrimSpace(tpl.RpcRouterTpl))
	if e != nil {
		panic(e.Error())
	}
	httpMap.MapHandlerMap = MapHandlerMap
	httpFile, _ := os.Create(RpcRouterFile)
	if err := tmpl.Execute(httpFile, httpMap); err != nil {
		panic(err.Error())
	}
	fmt.Println("http 路由文件生成成功")
}
