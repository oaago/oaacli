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

var GenType = &cobra.Command{
	Use:   "types",
	Short: "oaacli types 生成文件",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 ./oaacli defined get@/aaa/bbb s")
			return
		}
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

		var ss = strings.Split(strings.ToLower(arg[1]), "/")
		fmt.Println(ss, "ss")
		dirName := ss[0]
		servicePath := "./internal/service/"
		ValidDefined(utils.Camel2Case(servicePath) + utils.Camel2Case(dirName))
		exists, err := utils.PathExists(utils.Camel2Case(servicePath) + utils.Camel2Case(dirName) + "/types.go")
		if err != nil || exists {
			fmt.Println("已存在文件无法生成", utils.Camel2Case(servicePath)+utils.Camel2Case(dirName)+"/types.go")
			return
		}
		fmt.Println("开始生成目录", utils.Camel2Case(servicePath)+utils.Camel2Case(dirName)+"/types.go")
		genType(servicePath, utils.Camel2Case(dirName), arg[0], ss[1])
	},
}

func genType(servicePath, dirName, method, fun string) {
	//模板变量
	type Defined struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Func      string
	}
	data := Defined{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    utils.Lcfirst(method),
		Func:      utils.Ucfirst(fun),
		UpMethod:  utils.Ucfirst(method),
	}
	//创建模板
	defined := "types"
	tmpl := template.New(defined)
	//解析模板
	text := tpl2.TYPESTPL
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err.Error())
	}
	typesDir := utils.Camel2Case(servicePath) + utils.Camel2Case(dirName)
	hasDir, _ := utils.PathExists(typesDir)
	if !hasDir {
		err := os.Mkdir(typesDir, os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	hasDir1, _ := utils.PathExists(typesDir + "/" + utils.Camel2Case(method))
	if !hasDir1 {
		e := os.Mkdir(typesDir + "/" + utils.Camel2Case(method), os.ModePerm)
		if e != nil {
			panic("目录初始化失败" + e.Error())
		}
	}
	//渲染输出
	fs, e := os.Create(typesDir + "/" + utils.Camel2Case(method) + "/types.go")
	if e != nil {
		fmt.Println("type 文件写入失败" + e.Error())
	}
	fmt.Println(typesDir, utils.Camel2Case(method) + "/types.go")
	tplerr := tpl.ExecuteTemplate(fs, defined, data)
	if tplerr != nil {
		panic(tplerr.Error())
	}
	fs.Close()
	fmt.Println("写入types模版成功 " + typesDir)
}

func genRpcType(servicePath, dirName, method, fun string) {
	//模板变量
	type Defined struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Func      string
	}
	data := Defined{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    utils.Lcfirst(method),
		Func:      utils.Ucfirst(fun),
		UpMethod:  utils.Ucfirst(method),
	}
	//创建模板
	defined := "defined"
	tmpl := template.New(defined)
	//解析模板
	text := tpl2.RpcTYPESTPL
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err.Error())
	}
	typesDir := utils.Camel2Case(servicePath) + utils.Camel2Case(dirName)
	hasDir, _ := utils.PathExists(typesDir)
	if !hasDir {
		err := os.Mkdir(typesDir, os.ModePerm)
		err = os.Mkdir(typesDir+"/"+utils.Camel2Case(method), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	//渲染输出
	fs, _ := os.Create(typesDir + "/" + utils.Camel2Case(method) + "/types.go")
	err = tpl.ExecuteTemplate(fs, defined, data)
	if err != nil {
		panic(err.Error())
	}
	fs.Close()
	fmt.Println("写入types模版成功 " + typesDir)
}

func ValidDefined(dirName string) {
	hasDir, _ := utils.PathExists(dirName)
	if !hasDir {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
}
