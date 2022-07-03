package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var GenInit = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "oaacli init 根据 oaago.json 生成出来需要的项目文件， 可以制定配置文件oaago.json 别名 i 例如 oaa i",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		uri := "./oaa.json"
		if len(args) == 1 {
			uri = args[0]
		}
		genDef(uri)
	},
}

func genDef(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	var mapurl map[string][]string
	json.Unmarshal(data, &mapurl)
	for _, team := range mapurl {
		for _, lis := range team {
			httpReg := regexp.MustCompile(`(get|post|put)@/[A-Za-z0-9]{0,20}/[A-Za-z0-9]{0,20}`)
			rpcReg := regexp.MustCompile(`(get|post|put)&/[A-Za-z0-9]{0,20}/[A-Za-z0-9]{0,20}`)
			lim := strings.Split(lis, "|")
			li := lim[0]
			result1 := httpReg.FindAllStringSubmatch(li, -1)
			result2 := rpcReg.FindAllStringSubmatch(li, -1)
			if len(result1) == 0 && len(result2) == 0 {
				panic(li + " 不符合规范请检查之后在使用")
			}
			apifilepath := "./internal/api/http/"
			rpcfileePath := "./internal/api/rpc/"
			routerPath := "./internal/router/"
			middlewarePath := "./internal/middleware/"
			servicePath := "./internal/service/"
			daoPath := "./internal/dao/"
			os.Mkdir(apifilepath, os.ModePerm)
			os.Mkdir(servicePath, os.ModePerm)
			os.Mkdir(routerPath, os.ModePerm)
			os.Mkdir(rpcfileePath, os.ModePerm)
			os.Mkdir(middlewarePath, os.ModePerm)
			os.Mkdir(daoPath, os.ModePerm)
			if strings.Contains(li, "@/") {
				arg := strings.Split(li, "@/")
				method := arg[0]
				str := arg[1]
				handlerStr := strings.Split(str, "/")
				fmt.Println(handlerStr, method)
				genType(servicePath, handlerStr[0], handlerStr[1], handlerStr[1])
				genApi(apifilepath, handlerStr[0], handlerStr[1], handlerStr[1], method)
				module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
				genRouter(module, utils.Ucfirst(handlerStr[0])+utils.Ucfirst(method), handlerStr[0], Url)
			} else if strings.Contains(li, "&/") {
				arg := strings.Split(li, "&/")
				ag := strings.Split(arg[1], "/")
				str := []string{ag[0] + "/" + ag[1]}
				genProto(str, "")
				genRpc(ProjectUrl+"internal/api/rpc/"+ag[0]+"/"+ag[1], ag[0], ag[0]+"_"+ag[1], ag[1])
				fmt.Println("proto 编译完成")
				genRpcServer(utils.Camel2Case("rpc_"+ag[0]), ag[1], ag[1], ag[0])
				fmt.Println("proto service 生成完成")
				module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
				genRpcRouter(module, utils.Ucfirst(ag[0])+utils.Ucfirst(ag[1]), ag[0], Url)
			} else {
				panic("不符合规范 http get@/aa/bb  rpc get&/aa/bb")
			}
		}
	}
}
