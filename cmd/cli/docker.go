package cli

import (
	"fmt"
	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/spf13/viper"
	"os"
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
			Port      int64
			Module    string
			HarborUrl string
		}
		Module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
		Op := viper.New()
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		Op.AddConfigPath(path)   //设置读取的文件路径
		Op.SetConfigName("app")  //设置读取的文件名
		Op.SetConfigType("yaml") //设置文件的类型
		Op.ReadInConfig()
		Port := Op.GetInt64("server.port")
		if Port == 0 {
			panic("docker 生成的配置文件必须有 port")
		}
		data := DockerFile{
			Module: Module,
			Port:   Port,
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
	Short: "打包成docker镜像 oaago dockerbuild v1.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		var version string
		Op := viper.New()
		path, _ := os.Getwd()
		Op.AddConfigPath(path)   //设置读取的文件路径
		Op.SetConfigName("app")  //设置读取的文件名
		Op.SetConfigType("yaml") //设置文件的类型
		Op.ReadInConfig()
		fmt.Println(Op.GetString("docker.harbor.url"), "harbor-url")
		if len(args) == 1 {
			version = args[0]
		} else {
			fmt.Println("请输入版本号，输入后即将打包(例如v1.0.0)：")
			fmt.Scanln(&version)
		}
		HarborUrl := strings.Replace(Op.GetString("docker.harbor.url"), "http", "", 1)
		module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
		fmt.Println("building....")
		//execCommand("docker", []string{"build", ".", "-t", HarborUrl + "/oaago/" + module + ":" + version, "-f", utils.GetCurrentPath() + "DockerFile"})
		out := utils.RunCmd("docker image ls |grep "+HarborUrl+"/oaago/"+module+":", true)
		ou := utils.RunCmd("docker build . -t "+HarborUrl+"/oaago/"+module+":"+version+" -f "+utils.GetCurrentPath()+"DockerFile", true)
		fmt.Println(string(ou))
		fmt.Println(string(out))
		var argStr string
		for _, arg := range args {
			argStr = argStr + arg
		}
		if strings.Contains(argStr, "-p") {
			pushOut := utils.RunCmd("docker push "+HarborUrl+"/oaago/"+module+":"+version, true)
			fmt.Println(string(pushOut))
		}
		fmt.Println("build end....")
	},
}
