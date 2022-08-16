package cli

import (
	"fmt"
	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"os"
	"os/exec"
	"strings"

	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var NewProject = &cobra.Command{
	Use:     "new",
	Aliases: []string{"new"},
	Short:   "示例 oaacli new project 生成项目",
	Run: func(cmd *cobra.Command, args []string) {
		if !check(args) {
			fmt.Println("命令行错误 请检查使用方式 示例 oaacgo new project")
			return
		}
		utils.CLIScreen()
		fmt.Println("开始生成项目........")
		ProjectUrl = currentPath + args[0]
		// 初始化目录
		initDir()
		// 初始化文件
		initFile(args[0])
		// 初始化环境
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
		fmt.Println(`项目初始化完成 开始更新依赖`)
		modOut := utils.RunCmd("go mod tidy", true)
		fmt.Println(string(modOut))
		fmt.Println("初始化完成")
	},
}

// 初始化必要的目录
func initDir() {
	fmt.Println("初始化目录")
	os.Mkdir(ProjectUrl, os.ModePerm) //nolint:errcheck
	os.Mkdir(ProjectUrl+"/internal", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/api", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/service", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/dao", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/model", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/router", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/consts", os.ModePerm)
	os.Mkdir(ProjectUrl+"/internal/middleware", os.ModePerm)
	if projectType == "a" {
		os.Mkdir(ProjectUrl+"/internal/api/http", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/middleware/http", os.ModePerm)
		os.Mkdir(routerPath, os.ModePerm)
		os.Mkdir(rpcfileePath, os.ModePerm)
		os.Mkdir(middlewarePath, os.ModePerm)
		os.Mkdir(daoPath, os.ModePerm)
	} else if projectType == "r" {
		os.Mkdir(ProjectUrl+"/internal/api/rpc", os.ModePerm)
		os.Mkdir(ProjectUrl+"/internal/middleware/rpc", os.ModePerm)
		os.Mkdir(ProjectUrl+"/rpc", os.ModePerm)
	}
}

// 初始化文件
func initFile(arg string) {
	fmt.Println("初始化文件")
	midFile, err := os.Create(ProjectUrl + "/internal/middleware/http/types.go")
	midFile.WriteString(tpl2.MiddlewareTpl)
	midFile.Close()

	codeFile, _ := os.Create(ProjectUrl + "/internal/consts/code.go")
	codeFile.WriteString(tpl2.ConstsTpl)
	codeFile.Close()

	definedFile, _ := os.Create(ProjectUrl + "/oaa.json")
	definedFile.WriteString(tpl2.OAATpl)
	definedFile.Close()

	errs := os.Chdir(arg)
	if errs != nil {
		return
	}
	mainFile, err := os.Create("main.go")
	if err != nil {
		panic(err.Error())
	}
	mainFile.WriteString(strings.Replace(tpl2.MainTpl, "%package%", arg, -1))
	mainFile.Close()
	ConfigFile, err := os.Create("app.yaml")
	if err != nil {
		panic(err.Error())
	}

	ConfigFile.WriteString(strings.Replace(tpl2.ConfingTpl, "%package%", arg, -1)) //nolint:errcheck
	ConfigFile.Close()
}

// 检查参数情况
func check(arg []string) bool {
	var specialCharacters = []string{"/", ",", ".", "。"}
	var result = true
	if len(arg) < 1 || len(arg) > 2 {
		result = false
	}
	if len(arg) == 2 {
		for _, s := range projectTypeMap {
			if "-"+arg[1] == s {
				projectType = s
			}
		}
		if projectType == "" {
			fmt.Println("使用参数 -a 代表api项目， -r 代表rpc项目")
		}
	}
	for _, character := range specialCharacters {
		if strings.Contains(arg[0], character) {
			result = false
			return false
		}
	}
	return result
}
