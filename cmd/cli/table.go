package cli

import (
	"fmt"
	"github.com/oaago/component/mysql"
	"github.com/oaago/component/mysql/gen"
	"github.com/spf13/cobra"
	"strings"
)

var GenTable = &cobra.Command{
	Use:   "table",
	Short: "示例 oaacli table scrm@t_user,user_base 在目录 ./internal/model 生成一个 model + query",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 oaacli table scrm@t_user,user_base")
			return
		}
		str := args[0]
		strCli := strings.Split(str, "@")
		dbName := ""
		table := make([]string, 0)
		if len(strCli) == 2 {
			dbName = strCli[0]
			table = strings.Split(strCli[1], ",")
		} else {
			dbName = "default"
			table = strings.Split(str, ",")
		}
		db, _ := mysql.NewConnect(dbName)
		if db == nil {
			panic(db.Error.Error() + "无法获取mysql")
		}
		if len(table) < 1 {
			fmt.Println("命令行错误  %s不正确", table)
			return
		}
		gen.GenStruct(dbName, table)
		for _, t := range table {
			genDao(dbName, t, t)
		}

	},
}
