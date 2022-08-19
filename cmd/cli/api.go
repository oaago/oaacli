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

// Api 模板变量
type Api struct {
	Package   string
	UpPackage string
	Method    string
	UpMethod  string
	Module    string
	//Met        []string
	//Param      map[string][]string
	//Upmet      []string
	//Dec        string
	//DecMessage map[string]string
	HandlerContent []HandlerContent
}

type HandlerContent struct {
	HandlerName string
	Method      string
	Param       []string
	Dec         string
	Comment     string
}

func genApi(apiPath, dirName, fileName, method, dec string, met []string) {
	// 验证目录是否存在
	hasDir, _ := utils.PathExists(utils.Camel2Case(apiPath) + utils.Camel2Case(dirName))
	if !hasDir {
		err := os.MkdirAll(utils.GetCurrentPath()+utils.Camel2Case(apiPath)+utils.Camel2Case(dirName), 0777)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
		e := os.Chmod(utils.GetCurrentPath()+utils.Camel2Case(apiPath)+utils.Camel2Case(dirName), 0777)
		if e != nil {
			return
		}
	}
	hasDir1, _ := utils.PathExists(utils.Camel2Case(apiPath) + utils.Camel2Case(dirName) + "/" + utils.Camel2Case(fileName))
	if !hasDir1 {
		err := os.MkdirAll(utils.GetCurrentPath()+utils.Camel2Case(apiPath)+utils.Camel2Case(dirName)+"/"+utils.Camel2Case(fileName), 0777)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
		e := os.Chmod(utils.GetCurrentPath()+utils.Camel2Case(apiPath)+utils.Camel2Case(dirName)+"/"+utils.Camel2Case(fileName), 0777)
		if e != nil {
			return
		}
	}
	// 根据types 获取所有文档参数
	types := utils.Camel2Case(servicePath) + utils.Camel2Case(dirName) + "/" + utils.Camel2Case(fileName)
	_, structList := utils.GetAllStruct(types)
	var param = make(map[string][]string)
	for s, tags := range structList {
		for _, tag := range tags {
			var paramName = tag.Name
			var paramType = tag.Type
			var contentType = "body"
			var required = "false"
			var comment = "-"
			for s2, s3 := range tag.Tags {
				if s2 == "json" {
					paramName = s3
				}
				if s2 == "validate" && strings.Contains(s3, "required") {
					required = "true"
				}
				if s2 == "comment" {
					comment = s3
				}
				if "Get"+utils.Ucfirst(dirName)+utils.Case2Camel(utils.Ucfirst(method))+"Req" == s {
					contentType = "query"
				}
			}
			str := "// @param " + paramName + " " + contentType + " " + paramType + " " + required + " \"" + comment + "\""
			fmt.Println("structName: " + s + " param: " + str + "\r\n")
			param[s] = append(param[s], str)
		}
	}
	// 请求类型转换
	HandlerContentList := make([]HandlerContent, 0)
	var Upmet = make([]string, 0)
	for _, s := range met {
		Upmet = append(Upmet, utils.Ucfirst(s))
		HandlerName := strings.Replace(SemanticMap[s], "$", utils.Ucfirst(dirName)+utils.Case2Camel(utils.Ucfirst(method)), 1)
		DecMsg := ""
		// 增加接口描述根据配置描述接口
		for k, msg := range DecMessage {
			if s == k {
				DecMsg = strings.Replace(msg, "$", dec, 1)
			}
		}
		HandlerContentList = append(HandlerContentList, HandlerContent{
			HandlerName: HandlerName + "Handler",
			Method:      utils.Ucfirst(s),
			Param:       param[HandlerName+"Req"],
			Dec:         DecMsg,
			Comment:     dec,
		})
	}
	data := Api{
		Package:        utils.Camel2Case(dirName),
		UpPackage:      utils.Ucfirst(dirName),
		Method:         utils.Lcfirst(method),
		UpMethod:       utils.Case2Camel(utils.Ucfirst(method)),
		Module:         module,
		HandlerContent: HandlerContentList,
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
	errs := os.MkdirAll(utils.Camel2Case(apiPath)+utils.Camel2Case(dirName)+"/"+utils.Camel2Case(fileName), 0777)
	if errs != nil {
		return
	}
	e := os.Chmod(utils.Camel2Case(apiPath)+utils.Camel2Case(dirName)+"/"+utils.Camel2Case(fileName), 0777)
	if e != nil {
		return
	}
	// 生成文件 渲染模版
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
