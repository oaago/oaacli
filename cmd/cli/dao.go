package cli

import (
	"fmt"
	"github.com/oaago/oaago/const"
	"os"
	"strings"
	"text/template"

	tpl "github.com/oaago/oaago/cmd/tpl"
	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var GenDao = &cobra.Command{
	Use:   "dao",
	Short: "oaacli dao name 根据name 生成dao",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 oaacli dao user(数据库配置连接名称)@sys_user(表)")
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
	//模板变量
	filesPath := strings.ToLower(utils.Camel2Case(_const.DaoPath+"dao_"+dirName) + "/" + fileName + ".go")
	exists, _ := utils.PathExists(filesPath)
	if exists {
		fmt.Println(_const.DaoPath + "dao_" + dirName + filesPath + "dao文件已经存在 不会继续创建")
		return
	}
	//模板变量
	type Dao struct {
		Package   string
		UpPackage string
		Method    string
		UpMethod  string
		Module    string
	}
	data := Dao{
		Package:   utils.Camel2Case(dirName),
		UpPackage: utils.Case2Camel(utils.Ucfirst(dirName)),
		Method:    utils.Lcfirst(method),
		UpMethod:  utils.Case2Camel(utils.Ucfirst(method)),
		Module:    _const.Module,
	}
	//创建模板
	daoDefind := "dao"
	tmpl := template.New(daoDefind)
	//解析模板
	text := tpl.DaoTpl
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err)
	}
	daoDir := utils.Camel2Case(_const.DaoPath) + utils.Camel2Case("dao_"+dirName)
	hasDir, _ := utils.PathExists(daoDir)
	if !hasDir {
		err := os.Mkdir(daoDir, os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	//渲染输出
	fs, _ := os.Create(filesPath)
	err = tpl.ExecuteTemplate(fs, daoDefind, data)
	fs.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(daoDir + filesPath + "写入dao模版成功 ")
}
