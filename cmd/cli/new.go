package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/oaago/oaacli/cmd/tpl"
	"github.com/oaago/oaacli/utils"
	"github.com/spf13/cobra"
)

var ProjectUrl = ""
var NewProject = &cobra.Command{
	Use:     "new",
	Aliases: []string{"new"},
	Short:   "示例 oaacli new project 生成项目",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 oaacli new project")
			return
		}
		fmt.Println("开始生成项目........")
		utils.CLIScreen()
		ProjectUrl = utils.GetCurrentPath() + args[0]
		os.Mkdir(ProjectUrl, os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/api", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/api/http", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/api/rpc", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/service", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/dao", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/dao/model", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/model", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/router", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/consts", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/middleware", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/middleware/http", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/middleware/rpc", os.ModePerm)
		midFile, err := os.Create(ProjectUrl + "/internal/middleware/http/types.go")
		midFile.WriteString(tpl.MIDTPL)
		midFile.Close()
		codeFile, _ := os.Create(ProjectUrl + "/internal/consts/code.go")
		codeFile.WriteString(tpl.CODETPL)
		codeFile.Close()
		definedFile, _ := os.Create(ProjectUrl + "/oaa.json")
		definedFile.WriteString(tpl.DEFIENDJSON)
		definedFile.Close()
		errs := os.Chdir(args[0])
		if errs != nil {
			return
		}
		mainFile, err := os.Create("main.go")
		if err != nil {
			panic(err.Error())
		}
		mainFile.WriteString(strings.Replace(tpl.MAINTPL, "%package%", args[0], -1))
		mainFile.Close()
		ConfigFile, err := os.Create("app.yaml")
		if err != nil {
			panic(err.Error())
		}
		ConfigFile.WriteString(strings.Replace(tpl.CONFIGTPL, "%package%", args[0], -1))
		ConfigFile.Close()
		exec.Command("bash", "-c", "cd ./"+args[0])
		fmt.Println("创建项目 " + args[0] + " go mod init ")
		output, _ := exec.Command("bash", "-c", "go mod init "+args[0]).Output()
		fmt.Println(string(output))
		fmt.Println("初始化 " + args[0] + "项目文件 oaacli init ")
		output, _ = exec.Command("bash", "-c", "oaacli init").Output()
		fmt.Println(string(output))
		fmt.Println("更新 " + args[0] + " 项目依赖   go mod tidy")
		output, _ = exec.Command("go", "mod tidy").Output()
		fmt.Println(string(output))
		fmt.Println(`项目初始化完成 请 更新依赖 go mod tidy`)
	},
}
