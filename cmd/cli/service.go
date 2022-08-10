package cli

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
)

// 根据http生成路径
func genServer(dirName, fileName, method string, met string) {
	// 检测是否存在types
	typePath := strings.ToLower(utils.Camel2Case(apiServicePath) + utils.Camel2Case(dirName) + "/" + "typs.go")
	exist, _ := utils.PathExists(typePath)
	if !exist {
		fmt.Println("types文件不存在 将会自动生成", typePath)
		genType(apiServicePath, utils.Camel2Case(dirName), fileName, utils.Lcfirst(method))
	}
	//模板变量
	filesPath := strings.ToLower(utils.Camel2Case(apiServicePath+dirName+"/"+fileName) + "/" + met + "_" + utils.Camel2Case(dirName) + "_" + utils.Lcfirst(method) + "_service.go")
	exists, _ := utils.PathExists(filesPath)
	if exists {
		fmt.Println("service文件已经存在 不会继续创建", filesPath)
		return
	}
	type Service struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Met       string
		Upmet     string
	}
	data := Service{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    method,
		UpMethod:  utils.Ucfirst(method),
		Met:       met,
		Upmet:     utils.Ucfirst(met),
	}
	//创建模板
	fmt.Println("开始写入service模版 " + fileName)
	service := "http-service"
	tmpl := template.New(service)
	//解析模板
	text := tpl.HttpServiceTpl
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err.Error() + "解析service模版失败")
	}
	//渲染输出
	hasDir, _ := utils.PathExists(utils.Camel2Case(apiServicePath) + utils.Camel2Case(dirName))
	if !hasDir {
		err := os.Mkdir(utils.Camel2Case(apiServicePath+dirName), os.ModePerm)
		err = os.Mkdir(utils.Camel2Case(apiServicePath+dirName+"/"+fileName), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	fmt.Println(filesPath, "httpfilesPath")
	fs, err := os.Create(filesPath)
	err = tpl.ExecuteTemplate(fs, service, data)
	if err != nil {
		panic(err.Error())
	}
	fs.Close()
	fmt.Println("写入http-service模版成功 " + filesPath)
}

func genRpcServer(dirName, fileName, method, dir string) {
	apiPath := "./internal/service/rpc/"
	// 检测是否存在types
	typePath := strings.ToLower(utils.Camel2Case(apiPath) + utils.Camel2Case(dirName) + "/" + fileName + "/types.go")
	exist, _ := utils.PathExists(typePath)
	if !exist {
		fmt.Println("types文件不存在 将会自动生成", typePath)
		genRpcType(apiPath, utils.Camel2Case(dirName), fileName, utils.Lcfirst(method))
	}
	//模板变量
	filesPath := strings.ToLower(utils.Camel2Case(apiPath+dirName+"/"+fileName) + "/" + utils.Camel2Case(dirName) + "_" + utils.Lcfirst(method) + "_service.go")
	exists, _ := utils.PathExists(filesPath)
	if exists {
		fmt.Println("service文件已经存在", filesPath)
		return
	}
	Module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	type Service struct {
		Package   string
		RpcName   string
		UpRpcName string
		UpPackage string
		Method    string
		UpMethod  string
		Module    string
	}
	data := Service{
		Package:   utils.Camel2Case(dirName),
		RpcName:   utils.Camel2Case(dir),
		UpRpcName: utils.Case2Camel(utils.Ucfirst(dir)),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    method,
		Module:    Module,
		UpMethod:  utils.Ucfirst(method),
	}
	//创建模板
	fmt.Println("开始写入rpc service模版 proto 名称" + fileName)
	tmpl := template.New("rpc-service")
	//解析模板
	tpl, err := tmpl.Parse(tpl.RpcServiceTpl)
	if err != nil {
		panic(err.Error() + "解析rpc service模版失败")
	}
	//渲染输出
	hasDir, _ := utils.PathExists(utils.Camel2Case(apiPath) + utils.Camel2Case(dirName))
	if !hasDir {
		err = os.Mkdir(utils.Camel2Case(apiPath+dirName), os.ModePerm)
		err := os.Mkdir(utils.Camel2Case(apiPath+dirName+"/"+fileName), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	fmt.Println(filesPath, "rpcfilesPath")
	fs, err := os.Create(filesPath)
	err = tpl.ExecuteTemplate(fs, "rpc-service", data)
	if err != nil {
		panic(err.Error())
	}
	fs.Close()
	fmt.Println("写入rpc-service模版成功 " + filesPath)
}
