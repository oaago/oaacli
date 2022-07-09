package cli

import (
	"bufio"
	"fmt"
	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var DockerFileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "自动生成 dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		type DockerFile struct {
			Port   int64
			Module string
		}
		module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
		Op := viper.New()
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		Op.AddConfigPath(path)   //设置读取的文件路径
		Op.SetConfigName("app")  //设置读取的文件名
		Op.SetConfigType("yaml") //设置文件的类型
		data := DockerFile{
			Module: module,
			Port:   Op.GetInt64("server.port"),
		}
		//创建模板
		dockerfile := "dockerfile"
		tmpl := template.New(dockerfile)
		//解析模板
		text := tpl2.DockerFile
		tpl, err := tmpl.Parse(text)
		if err != nil {
			panic(err.Error())
		}
		//渲染输出
		utils.RunCmd("rm -rf ./DockerFile", true)
		fs, _ := os.Create("DockerFile")
		err = tpl.ExecuteTemplate(fs, dockerfile, data)
		if err != nil {
			panic(err.Error())
		}
		fs.Close()
		fmt.Println("写入dockerfile模版成功 ")
	},
}

var DockerBuildCmd = &cobra.Command{
	Use:   "dockerbuild",
	Short: "打包成docker镜像",
	Run: func(cmd *cobra.Command, args []string) {
		var version string
		fmt.Println("请输入版本号，输入后即将打包(例如v1.0.0)：")
		fmt.Scanln(&version)
		module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
		execCommand("docker", []string{"build", ".", "-t", "oaago/" + module + ":" + version, "-f", utils.GetCurrentPath() + "DockerFile"})
	},
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
