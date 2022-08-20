package _const

type GenFuncMap struct {
	Method       string
	FunctionName string
}

var (
	ApiServicePath = "./internal/service/"
	DaoPath        = "./internal/dao/"
	Module         = ""
	ProjectUrl     = ""
	CurrentPath    = ""
	ProjectTypeMap = []string{"a", "r"} //a 代表api r代表rpc
	ProjectType    = "a"
	ConfigFile     = "./oaa.json"
	ServicePath    = "./internal/service/"
	Apifilepath    = "./internal/api/"
	RpcfileePath   = "./internal/api/rpc/"
	RouterPath     = "./internal/router/"
	MiddlewarePath = "./internal/middleware/"
	HttpRouterFile = "./internal/router/router-http.go"
	RpcRouterFile  = "./internal/router/router-http.go"
	AllowMethods   = "get,post,delete,put,patch"
	SemanticMap    = make([]GenFuncMap, 0)
	DecMessage     = map[string]string{
		"get":    "获取$信息",
		"put":    "更新$信息",
		"post":   "创建$信息",
		"delete": "删除$信息",
		"patch":  "更新$相关字段",
	}
	TableMap = map[string][]string{}
)

func init() {
	SemanticMap = append(SemanticMap, GenFuncMap{
		Method:       "get",
		FunctionName: "Get$Info",
	}, GenFuncMap{
		Method:       "get",
		FunctionName: "Get$InfoList",
	}, GenFuncMap{
		Method:       "post",
		FunctionName: "Create$",
	}, GenFuncMap{
		Method:       "put",
		FunctionName: "Update$ById",
	}, GenFuncMap{
		Method:       "put",
		FunctionName: "Update$StatusById",
	}, GenFuncMap{
		Method:       "delete",
		FunctionName: "Delete$ByIds",
	}, GenFuncMap{
		Method:       "patch",
		FunctionName: "UpdateField$ByIds",
	})
}
