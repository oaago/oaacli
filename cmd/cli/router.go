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
	HasMid              bool
}

func genRouter(module, pack string) {
	MapHandlerMap := make([]MapHttpHandler, 0)
	MapHandlerMapImport := make([]MapHttpHandler, 0)
	HandlerMapOfOne := map[string]bool{}
	var httpMap HttpMap
	var HasMid = false //是否需要中间件import
	data, _ := os.ReadFile(configFile)
	httpData := gjson.Get(string(data), "http").Array()
	for _, datum := range httpData {
		lis := strings.Split(datum.String(), "@/")
		if len(lis) != 2 {
			panic(datum.Str + "不符合规则")
		}
		// 去除备注
		decStr := strings.Split(lis[1], "**")
		if len(decStr) == 2 {
			lis[1] = strings.Replace(lis[1], "**"+decStr[1], "", 1)
		}
		// 如果是 * 代表支持所以请求方式
		if lis[0] == "*" {
			lis[0] = strings.Replace(lis[0], "*", AllowMethods, 1)
		}
		lim := strings.Split(lis[1], "|")
		dir := lim[0]
		hand := strings.Split(dir, "/")
		if len(hand) > 2 {
			hand1 := strings.Replace(dir, "/", "_", -1)
			hand = []string{hand[0], strings.Replace(hand1, hand[0]+"_", "", 1)}
		}
		Middleware := make([]string, 0)
		if len(lim) > 1 {
			HasMid = true
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
			if HandlerMapOfOne[hand[0]+hand[1]] != true {
				HandlerMapOfOne[hand[0]+hand[1]] = true
				MapHandlerMapImport = append(MapHandlerMapImport, MapHttpHandler{
					Url:             s + "@/" + dir,
					Module:          module,
					Middleware:      Middleware,
					Handler:         utils.Case2Camel(utils.Ucfirst(hand[0]) + utils.Ucfirst(hand[1])),
					HttpDir:         utils.Case2Camel(utils.Camel2Case(hand[0])),
					Method:          hand[1],
					UpMethod:        utils.Case2Camel(utils.Ucfirst(hand[1])),
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
				Handler:    utils.Case2Camel(utils.Ucfirst(hand[0]) + utils.Ucfirst(hand[1])),
				HttpDir:    utils.Case2Camel(utils.Camel2Case(hand[0])),
				Method:     hand[1],
				UpMethod:   utils.Case2Camel(utils.Ucfirst(hand[1])),
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
	httpMap.HasMid = HasMid
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
