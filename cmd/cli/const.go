package cli

import (
	"github.com/oaago/oaago/utils"
	"strings"
)

var (
	apiServicePath = "./internal/service/"
	daoPath        = "./internal/dao/"
	module         = strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	ProjectUrl     = ""
	currentPath    = utils.GetCurrentPath()
	projectTypeMap = []string{"a", "r"} //a 代表api r代表rpc
	projectType    = "a"
	configFile     = "./oaa.json"
	servicePath    = "./internal/service/"
	apifilepath    = "./internal/api/"
	rpcfileePath   = "./internal/api/rpc/"
	routerPath     = "./internal/router/"
	middlewarePath = "./internal/middleware/"
	HttpRouterFile = "./internal/router/router-http.go"
	RpcRouterFile  = "./internal/router/router-http.go"
	AllowMethods   = "get,post,delete,put,patch"
	SemanticMap    = map[string]string{
		"get":    "Get$ById",
		"post":   "Create$",
		"delete": "Delete$ById",
		"put":    "Update$",
		"patch":  "Update$",
	}
	DecMessage = map[string]string{
		"get":    "获取$信息",
		"post":   "创建$信息",
		"delete": "删除$信息",
		"put":    "更新$信息",
		"patch":  "更新$相关字段",
	}
	TableMap = map[string][]string{}
)
