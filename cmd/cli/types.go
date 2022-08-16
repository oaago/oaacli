package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
)

func genType(servicePath, dirName, method, fun, currentDBName string) {
	//模板变量
	var met = make([]string, 0)
	for s, _ := range DecMessage {
		met = append(met, utils.Ucfirst(s)+utils.Case2Camel(utils.Ucfirst(dirName))+utils.Case2Camel(utils.Ucfirst(method)))
	}
	type Defined struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Func      string
		Met       []string
		DBName    string
		Module    string
	}
	data := Defined{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    utils.Lcfirst(method),
		Func:      utils.Ucfirst(fun),
		UpMethod:  utils.Case2Camel(utils.Ucfirst(method)),
		Met:       met,
		DBName:    currentDBName,
		Module:    module,
	}
	//创建模板
	defined := "types"
	tmpl := template.New(defined)
	//解析模板
	text := tpl2.HttpTypesTpl
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err)
	}
	typesDir := utils.Camel2Case(servicePath) + utils.Camel2Case(dirName)
	//渲染输出
	hasFile, _ := utils.PathExists(typesDir + "/" + utils.Camel2Case(method) + "/service.go")
	if hasFile {
		fmt.Println(typesDir + "/" + utils.Camel2Case(method) + "/service.go" + "文件已存在，不会继续创建")
	}
	fs, e := os.Create(typesDir + "/" + utils.Camel2Case(method) + "/service.go")
	if e != nil {
		fmt.Println("type 文件写入失败" + e.Error())
	}
	tplerr := tpl.ExecuteTemplate(fs, defined, data)
	if tplerr != nil {
		panic(tplerr)
	}
	fs.Close()
	cmd := exec.Command("gofmt", "-w", typesDir+"/"+utils.Camel2Case(method)+"/service.go")
	cmd.Run() //nolint:errcheck
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
