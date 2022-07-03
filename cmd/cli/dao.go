package cli

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	tpl2 "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var GenDao = &cobra.Command{
	Use:   "dao",
	Short: "oaacli dao name 根据name 生成dao",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 ./oaacli dao /aaa/bbb")
			return
		}
		var ss = strings.Split(args[0], "@")
		fmt.Println(args[0])
		dirName := ss[0]
		fileName := ss[1]
		fmt.Println("开始生成dao", "目录是"+dirName, "函数名称是:", dirName+fileName+"service")
		genDao(dirName, fileName, "service")
	},
}

func genDao(dirName, fileName string, method string) {
	daoPath := "./internal/dao/"
	//模板变量
	filesPath := strings.ToLower(utils.Camel2Case(daoPath+dirName) + "/" + fileName + ".go")
	exists, _ := utils.PathExists(filesPath)
	if exists {
		fmt.Println("service文件已经存在 不会继续创建", filesPath)
		return
	}
	//模板变量
	type Defined struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Module    string
	}
	module := strings.Replace(string(utils.RunCmd("go list -m", true)), "\n", "", -1)
	data := Defined{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    utils.Lcfirst(method),
		UpMethod:  utils.Case2Camel(utils.Ucfirst(method)),
		Module:    module,
	}
	//创建模板
	defined := "dao"
	tmpl := template.New(defined)
	//解析模板
	text := tpl2.DAOTPL
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err.Error())
	}
	daoDir := utils.Camel2Case(daoPath) + utils.Camel2Case(dirName)
	hasDir, _ := utils.PathExists(daoDir)
	if !hasDir {
		err := os.Mkdir(daoDir, os.ModePerm)
		//err = os.Mkdir(daoDir+"/"+utils.Camel2Case(method), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	//渲染输出
	fs, _ := os.Create(filesPath)
	err = tpl.ExecuteTemplate(fs, defined, data)
	if err != nil {
		panic(err.Error())
	}
	fs.Close()
	fmt.Println("写入dao模版成功 " + filesPath)
}
