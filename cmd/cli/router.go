package cli

import (
	"encoding/json"
	"fmt"
	_const2 "github.com/oaago/oaago/const"
	"os"
	"os/exec"
	"strings"
	"text/template"

	tpl "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
	"github.com/tidwall/gjson"
)

type MapHttpHandler struct {
	Url             string
	RequestUrl      string
	RequestType     string
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
	RN              bool
}

type HttpMap struct {
	MapHandlerMap       []MapHttpHandler
	MiddlewareLen       int
	Module              string
	MapHandlerMapImport []MapHttpHandler
	HasMid              bool
}

type Props struct {
	ModuleName    string              `json:"moduleName,omitempty"`
	Url           string              `json:"url,omitempty"`
	Method        string              `json:"method,omitempty"`
	FieldName     string              `json:"fieldName,omitempty"`
	Description   string              `json:"description,omitempty"`
	Type          string              `json:"type,omitempty"`
	Validate      string              `json:"validate,omitempty"`
	ValidateRules map[string][]string `json:"validateRules"`
	TableName     string              `json:"tableName,omitempty"`
}

func genRouter(module, pack string) {
	MapHandlerMap := make([]MapHttpHandler, 0)
	MapHandlerMapImport := make([]MapHttpHandler, 0)
	HandlerMapOfOne := map[string]bool{}
	var httpMap HttpMap
	var HasMid = false //是否需要中间件import
	data, _ := os.ReadFile(_const2.ConfigFile)
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
			lis[0] = strings.Replace(lis[0], "*", _const2.AllowMethods, 1)
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
		methodMap := make([]string, 0)
		if !strings.Contains(lis[0], ",") {
			methodMap = []string{lis[0]}
		} else {
			methodMap = strings.Split(lis[0], ",")
		}
		for i, s := range methodMap {
			// 判断有多少个handler 文件
			if HandlerMapOfOne[hand[0]+hand[1]] != true {
				HandlerMapOfOne[hand[0]+hand[1]] = true
				MapHandlerMapImport = append(MapHandlerMapImport, MapHttpHandler{
					Module:          module,
					HttpDir:         utils.Case2Camel(utils.Camel2Case(hand[0])),
					Method:          utils.Camel2Case(utils.Lcfirst(hand[1])),
					UpMethod:        utils.Case2Camel(utils.Ucfirst(hand[1])),
					HandlerMapOfOne: HandlerMapOfOne,
				})
			}
			for _, funcMap := range _const2.SemanticMap {
				if strings.ToLower(s) == strings.ToLower(funcMap.Method) {
					//urlPk := strings.Replace(funcMap.FunctionName, "$", "", 1)
					HandlerName := strings.Replace(funcMap.FunctionName, "$", utils.Ucfirst(hand[0])+utils.Case2Camel(utils.Ucfirst(hand[1])), 1)
					MapHandlerMap = append(MapHandlerMap, MapHttpHandler{
						Url:         s + "@/" + dir,
						RequestUrl:  "/" + dir + "/" + HandlerName,
						RequestType: strings.ToUpper(s),
						Module:      module,
						Middleware:  Middleware,
						Handler:     HandlerName + "Handler",
						HttpDir:     utils.Case2Camel(utils.Camel2Case(hand[0])),
						Method:      hand[1],
						UpMethod:    utils.Case2Camel(utils.Ucfirst(hand[1])),
						Package:     hand[1],
						UpPackage:   utils.Camel2Case(pack),
						Upmet:       utils.Ucfirst(s),
						RN:          i == 0,
					})
				}
			}
		}
		if len(Middleware) != 0 {
			httpMap.MiddlewareLen = len(Middleware)
		}
	}
	apiData := gjson.Get(string(data), "api").Array()
	for _, datum := range apiData {
		var li Props
		err := json.Unmarshal([]byte(datum.String()), &li)
		if err != nil {
			return
		}
		if li.Url[0] == '/' {
			li.Url = strings.Replace(li.Url, "/", "", 1)
		}
		hand := strings.Split(li.Url, "/")
		if HandlerMapOfOne[hand[0]+hand[1]] != true {
			HandlerMapOfOne[hand[0]+hand[1]] = true
			MapHandlerMapImport = append(MapHandlerMapImport, MapHttpHandler{
				Module:          module,
				HttpDir:         utils.Case2Camel(utils.Camel2Case(hand[0])),
				Method:          utils.Camel2Case(utils.Lcfirst(hand[1])),
				UpMethod:        utils.Case2Camel(utils.Ucfirst(hand[1])),
				HandlerMapOfOne: HandlerMapOfOne,
			})
		}
		for _, funcMap := range _const2.SemanticMap {
			if strings.ToLower(li.Method) == strings.ToLower(funcMap.Method) {
				HandlerName := strings.Replace(funcMap.FunctionName, "$", utils.Ucfirst(hand[0])+utils.Case2Camel(utils.Ucfirst(hand[1])), 1)
				MapHandlerMap = append(MapHandlerMap, MapHttpHandler{
					Url:         li.Method + "@/" + li.Url,
					RequestUrl:  "/" + li.Url + "/" + HandlerName,
					RequestType: strings.ToUpper(li.Method),
					Module:      module,
					Middleware:  []string{},
					Handler:     HandlerName + "Handler",
					HttpDir:     utils.Case2Camel(utils.Camel2Case(hand[0])),
					Method:      hand[1],
					UpMethod:    utils.Case2Camel(utils.Ucfirst(hand[1])),
					Package:     hand[1],
					UpPackage:   utils.Camel2Case(pack),
					Upmet:       utils.Ucfirst(li.Method),
					RN:          true,
				})
			}
		}
	}
	httpMap.MapHandlerMap = MapHandlerMap
	httpMap.Module = module
	httpMap.MapHandlerMapImport = MapHandlerMapImport
	httpMap.HasMid = HasMid
	err := os.MkdirAll(_const2.RouterPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err1 := os.Chmod(_const2.RouterPath, os.ModePerm)
	if err1 != nil {
		panic(err1)
	}
	httpFile, _ := os.Create(_const2.HttpRouterFile)
	tmpl, e := template.New("gen-http-router").Parse(strings.TrimSpace(tpl.HttpRouterTpl))
	if e != nil {
		panic(e.Error())
	}
	if err2 := tmpl.Execute(httpFile, httpMap); err != nil {
		panic(err2)
	}
	cmd := exec.Command("gofmt", "-w", _const2.HttpRouterFile)
	cmd.Run() //nolint:errcheck
	fmt.Println("http 路由文件生成成功")
}
func genRpcRouter(module, handler, pack, url string) {
	MapHandlerMap := make([]MapHttpHandler, 0)
	var httpMap HttpMap
	data, _ := os.ReadFile(_const2.ConfigFile)
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
			Handler:   utils.Ucfirst(hand[0]) + utils.Ucfirst(hand[1]) + "Handler",
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
	httpFile, _ := os.Create(_const2.RpcRouterFile)
	if err := tmpl.Execute(httpFile, httpMap); err != nil {
		panic(err.Error())
	}
	fmt.Println("http 路由文件生成成功")
}
