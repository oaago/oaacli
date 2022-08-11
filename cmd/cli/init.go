package cli

import (
	"encoding/json"
	"fmt"
	tpl "github.com/oaago/oaago/cmd/tpl"
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
		if len(args) == 0 {
			genDef()
		}
	},
}

func genDef() {
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	var mapurl map[string][]string
	json.Unmarshal(data, &mapurl)
	// 初始化目录
	initFile(module)
	hasRpc := false
	for _, team := range mapurl {
		for _, lis := range team {
			// 先验证规则是否合法
			httpReg := regexp.MustCompile(`(get|post|put|delete|update|\*)@(/[A-Za-z0-9]{0,20})+\*\*.*`)
			rpcReg := regexp.MustCompile(`(get|post|put|delete|update|\*)&(/[A-Za-z0-9]{0,20})+\*\*.*`)
			lim := strings.Split(strings.ToLower(lis), "|")
			li := lim[0]
			result1 := httpReg.FindAllStringSubmatch(li, -1)
			result2 := rpcReg.FindAllStringSubmatch(li, -1)
			if len(result1) == 0 && len(result2) == 0 {
				panic(li + " 不符合规范请检查之后在使用")
			}
			// 解析接口描述
			// 根据规则获取不同的参数
			// 1. @ 代表的是http通信 请求模式@请求地址
			// 2. & 代表的是rpc通信
			// 3. | 代表的是中间件
			// 4. ** 代表的是 备注 对于接口的描述
			var dec string
			decStr := strings.Split(li, "**")
			if len(decStr) == 2 {
				// 对于接口的描述
				dec = decStr[1]
				li = strings.Replace(li, "**"+decStr[1], "", 1)
			}
			// 解析目录结构
			// handlerStr 0 代表目录 1代表文件 0_1 代表名称
			if strings.Contains(li, "@/") {
				arg := strings.Split(li, "@/")
				if arg[0] == "*" {
					arg[0] = "get,post,delete,put,update"
				}
				str := arg[1]
				// 解析模版
				handlerStr := strings.Split(str, "/")
				fmt.Println(handlerStr, "路由信息")
				if len(handlerStr) > 2 {
					hand1 := strings.Replace(str, "/", "_", -1)
					handlerStr = []string{handlerStr[0], strings.Replace(hand1, handlerStr[0]+"_", "", 1)}
				}
				genType(servicePath, handlerStr[0], handlerStr[1], handlerStr[1])
				// arg[0] 代表的是请求方法 arg[1] 请求路径
				mothedMap := strings.Split(arg[0], ",")
				for _, s := range mothedMap {
					has := strings.Contains(AllowMethods, s)
					if !has {
						fmt.Printf("检测出请求方式" + arg[0] + "存在" + s + "不正确 没有对应的 method\n")
						return
					}
				}
				genApi(apifilepath, handlerStr[0], handlerStr[1], handlerStr[1], dec, mothedMap)
				fmt.Println("开始装载路由...." + utils.Camel2Case(handlerStr[0]) + handlerStr[1])
				genRouter(module, handlerStr[0])
				fmt.Println("http初始化成功！")
			} else if strings.Contains(li, "&/") {
				hasRpc = true
				arg := strings.Split(li, "&/")
				ag := strings.Split(arg[1], "/")
				str := []string{ag[0] + "/" + ag[1]}
				genProto(str, "")
				fmt.Println("proto 编译完成")
				genRpcServer(utils.Camel2Case(ag[0]), ag[1], ag[1], ag[0])
				fmt.Println("proto service 生成完成")
				module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
				genRpcRouter(module, utils.Ucfirst(ag[0])+utils.Ucfirst(ag[1]), ag[0], arg[1])
				_, err := os.Stat("./powerproto.yaml")
				if err != nil {
					pow, _ := os.Create("./powerproto.yaml")
					pow.WriteString(tpl.PowerprotoTpl)
					pow.Close()
				}
				fmt.Println("rpc初始化成功！")
			} else {
				panic("不符合规范 http get@/aa/bb  rpc get&/aa/bb")
			}
		}
	}
	if hasRpc {
		mainFile, err := os.Create("main.go")
		if err != nil {
			panic(err.Error())
		}
		// 处理包名称
		def := strings.Replace(tpl.MainTpl, "%package%", module, -1)
		// 处理是否增加rpc server
		newTpl := strings.Replace(def, "//route.RpcServer", "route.RpcServer", 1)
		mainFile.WriteString(newTpl)
		mainFile.Close()
		fmt.Println("新增rpc处理模式")
	}
	modOut := utils.RunCmd("go mod tidy", true)
	fmt.Println(string(modOut))
	swagOut := utils.RunCmd("swag init", true)
	fmt.Println(string(swagOut))
	fmt.Println("初始化完成")
}
