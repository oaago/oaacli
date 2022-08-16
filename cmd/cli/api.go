package cli

import (
	_ "embed"
	"fmt"
	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/oaago/oaago/utils"
)

func genApi(apiPath, dirName, fileName, method, dec string, met []string) {
	var Upmet = []string{}
	for _, s := range met {
		Upmet = append(Upmet, utils.Ucfirst(s))
	}
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
		Package    string
		UpPackage  string
		Method     string
		UpMethod   string
		Module     string
		Met        []string
		Upmet      []string
		Dec        string
		DecMessage map[string]string
	}
	module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	var DecMsg = make(map[string]string)
	for _, mm := range met {
		for k, msg := range DecMessage {
			if mm == k {
				DecMsg[utils.Ucfirst(k)] = strings.Replace(msg, "$", dec, 1)
			}
		}
	}
	data := Api{
		Package:    utils.Camel2Case(dirName),
		UpPackage:  utils.Ucfirst(dirName),
		Method:     utils.Lcfirst(method),
		UpMethod:   utils.Case2Camel(utils.Ucfirst(method)),
		Module:     module,
		Met:        met,
		Upmet:      Upmet,
		Dec:        dec,
		DecMessage: DecMsg,
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
	//exists, err := utils.PathExists(strings.ToLower(filesName))
	//if err != nil || !exists {
	//	fmt.Println(filesName + "文件不存在即将创建文件")
	//	fs, _ := os.Create(filesName)
	//	err = tpl.ExecuteTemplate(fs, api, data)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fs.Close()
	//	fmt.Println(dirName + filesName + " api模版创建成功, 开始执行service 创建")
	//	for _, s := range met {
	//		s2 := SemanticMap[s]
	//		if len(s2) != 0 {
	//			genServer(utils.Camel2Case(dirName), fileName, fileName, s)
	//		}
	//	}
	//}
	//fmt.Println(filesName + "文件不存在即将创建文件")
	fs, _ := os.Create(filesName)
	err = tpl.ExecuteTemplate(fs, api, data)
	if err != nil {
		panic(err)
	}
	fs.Close()
	fmt.Println(dirName + filesName + " api模版创建成功, 开始执行service 创建")
	for _, s := range met {
		s2 := SemanticMap[s]
		if len(s2) != 0 {
			genServer(utils.Camel2Case(dirName), fileName, fileName, s)
		}
	}
	cmd := exec.Command("gofmt", "-w", filesName)
	cmd.Run() //nolint:errcheck
	fmt.Println("文件已存在不再生成" + filesName)
}
