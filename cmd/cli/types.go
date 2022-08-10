package cli

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
)

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
	text := tpl2.HttpTypesTpl
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
		e := os.Mkdir(typesDir+"/"+utils.Camel2Case(method), os.ModePerm)
		if e != nil {
			panic("目录初始化失败" + e.Error())
		}
	}
	//渲染输出
	fs, e := os.Create(typesDir + "/" + utils.Camel2Case(method) + "/types.go")
	if e != nil {
		fmt.Println("type 文件写入失败" + e.Error())
	}
	fmt.Println(typesDir, utils.Camel2Case(method)+"/types.go")
	tplerr := tpl.ExecuteTemplate(fs, defined, data)
	if tplerr != nil {
		panic(tplerr)
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
		Module    string
	}
	module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	data := Defined{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    utils.Lcfirst(method),
		Func:      utils.Ucfirst(fun),
		Module:    module,
		UpMethod:  utils.Ucfirst(method),
	}
	//创建模板
	defined := "rpctype"
	tmpl := template.New(defined)
	//解析模板
	text := tpl2.RpcTypesTpl
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err)
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
		panic(err)
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
