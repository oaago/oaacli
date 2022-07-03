package cli

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var Url = ""

var GenApi = &cobra.Command{
	Use:   "api",
	Short: "示例 oaacli api get@/staff/info s 生成一个api + service",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 ./oaacli api get@/aaa/bbb")
			return
		}
		Url = args[0]
		arg := strings.Split(args[0], "@/")
		if len(arg) < 2 {
			fmt.Println("命令行错误  %s不正确", arg[0])
			return
		}
		method := "get,post,delete,put,head,options"
		has := strings.Contains(method, arg[0])
		if !has {
			fmt.Printf("命令行错误  " + arg[0] + "不正确 没有对应的 method\n")
			return
		}
		var ss = strings.Split(arg[1], "/")
		dirName := ss[0]
		fileName := ss[0]
		if strings.Contains(args[0], "@/") {
			apiPath := "./internal/api/http/"
			genApi(apiPath, dirName, dirName, fileName, arg[0])
		} else if strings.Contains(args[0], "&/") {
			arg := strings.Split(args[0], "&/")
			genProto([]string{arg[1]}, "")
		}
	},
}

func genApi(apiPath, dirName, fileName, method, met string) {
	//fmt.Println(apiPath, dirName, fileName, method, "apiPath, dirName, fileName, method")
	hasDir, _ := utils.PathExists(utils.Camel2Case(apiPath) + utils.Camel2Case(dirName))
	if !hasDir {
		err := os.Mkdir(utils.GetCurrentPath()+utils.Camel2Case(apiPath)+utils.Camel2Case(dirName), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	hasDir1, _ := utils.PathExists(utils.Camel2Case(apiPath) + utils.Camel2Case(dirName) + "/" + utils.Camel2Case(fileName))
	if !hasDir1 {
		err := os.Mkdir(utils.GetCurrentPath()+utils.Camel2Case(apiPath)+utils.Camel2Case(dirName)+"/"+utils.Camel2Case(fileName), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	//模板变量
	type Api struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Module    string
		Met       string
	}
	module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	data := Api{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Ucfirst(dirName),
		Method:    utils.Lcfirst(method),
		UpMethod:  utils.Ucfirst(method),
		Module:    module,
		Met:       met,
	}
	//创建模板
	fmt.Println("开始api写入模版 " + fileName)
	api := "api"
	tmpl := template.New(api)
	//解析模板
	text := tpl2.ApiTPL
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err)
	}	
	//渲染输出
	filesName := utils.Camel2Case(apiPath) + utils.Camel2Case(dirName) + "/" + utils.Camel2Case(fileName) + "/" + utils.Camel2Case(dirName) + "_" + utils.Camel2Case(fileName) + "_handler.go"
	exists, err := utils.PathExists(strings.ToLower(filesName))
	if err != nil || exists {
		fmt.Println("文件不存在 即将创建文件")
	}
	fs, _ := os.Create(filesName)
	err = tpl.ExecuteTemplate(fs, api, data)
	if err != nil {
		panic(err.Error())
	}
	fs.Close()
	fmt.Println("写入api模版成功 " + filesName)
	genServer(utils.Camel2Case(dirName), fileName, fileName)
	fmt.Println("开始装载路由...." + Url)
}
