package cli

import (
	"fmt"
	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/const"
	"os"
	"os/exec"
	"strings"

	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var NewProject = &cobra.Command{
	Use:     "new",
	Aliases: []string{"new"},
	Short:   "示例 oaago new project 生成项目",
	Run: func(cmd *cobra.Command, args []string) {
		if !check(args) {
			fmt.Println("命令行错误 请检查使用方式 示例 oaacgo new project")
			return
		}
		utils.CLIScreen()
		fmt.Println("开始生成项目........")
		_const.ProjectUrl = _const.CurrentPath + args[0]
		// 初始化目录
		initDir()
		// 初始化文件
		initFile(args[0])
		// 初始化环境
		exec.Command("bash", "-c", "cd ./"+args[0])
		fmt.Println("创建项目 " + args[0] + " go mod init ")
		output, _ := exec.Command("bash", "-c", "go mod init "+args[0]).Output()
		fmt.Println(string(output))
		fmt.Println("初始化 " + args[0] + "项目文件 oaago init ")
		output, _ = exec.Command("bash", "-c", "oaago init").Output()
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
	os.Mkdir(_const.ProjectUrl, os.ModePerm) //nolint:errcheck
	_ = os.Mkdir(_const.ProjectUrl+"/internal", os.ModePerm)
	os.MkdirAll(_const.ProjectUrl+_const.Apifilepath, os.ModePerm)
	os.MkdirAll(_const.ProjectUrl+_const.ServicePath, os.ModePerm)
	os.MkdirAll(_const.ProjectUrl+_const.DaoPath, os.ModePerm)
	os.MkdirAll(_const.ProjectUrl+_const.RouterPath, os.ModePerm)
	os.MkdirAll(_const.ProjectUrl+_const.ConstPath, os.ModePerm)
	os.MkdirAll(_const.ProjectUrl+_const.MiddlewarePath, os.ModePerm)
	if _const.ProjectType == "a" {
		os.MkdirAll(_const.ProjectUrl+_const.Apifilepath, os.ModePerm)
		os.MkdirAll(_const.ProjectUrl+_const.MiddlewareHttpPath, os.ModePerm)
		os.MkdirAll(_const.RouterPath, os.ModePerm)
		os.MkdirAll(_const.MiddlewarePath, os.ModePerm)
		os.MkdirAll(_const.DaoPath, os.ModePerm)
	} else if _const.ProjectType == "r" {
		os.MkdirAll(_const.ProjectUrl+_const.RpcfileePath, os.ModePerm)
		os.MkdirAll(_const.ProjectUrl+_const.MiddlewareRpcPath, os.ModePerm)
		os.MkdirAll(_const.ProjectUrl+"/rpc", os.ModePerm)
	}
}

// 初始化文件
func initFile(arg string) {
	fmt.Println("初始化文件")
	midFile, err := os.Create(_const.ProjectUrl + _const.MiddlewareHttpPath + "/types.go")
	midFile.WriteString(tpl2.MiddlewareTpl)
	midFile.Close()

	codeFile, _ := os.Create(_const.ProjectUrl + _const.ConstPath + "/code.go")
	codeFile.WriteString(tpl2.ConstsTpl)
	codeFile.Close()

	definedFile, _ := os.Create(_const.ProjectUrl + "/oaa.json")
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
		for _, s := range _const.ProjectTypeMap {
			if "-"+arg[1] == s {
				_const.ProjectType = s
			}
		}
		if _const.ProjectType == "" {
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
