package cli

import (
	"fmt"
	"github.com/oaago/cloud/mysql"
	"github.com/oaago/cloud/mysql/gen"
	"github.com/spf13/cobra"
	"strings"
)

var GenTable = &cobra.Command{
	Use:   "table",
	Short: "示例 oaago table scrm@user 在目录 ./internal/model 生成一个 model + query",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("命令行错误 请检查使用方式 示例 oaacli table scrm@t_user,user_base")
			return
		}
		str := args[0]
		strCli := strings.Split(str, "@")
		dbName := strCli[0]
		table := make([]string, 0)
		if len(strCli) == 1 {
			strCli = append(strCli, "-")
		}
		if len(strCli) == 2 {
			if strCli[1] == "" {
				strCli[1] = "-"
			}
			if strCli[1] == "-" {
				GetTables(dbName)
				return
			}
			table = strings.Split(strCli[1], ",")
		} else {
			panic("请配置数据库连接")
		}
		db, _ := mysql.NewConnect(dbName)
		if db == nil {
			panic(db.Error.Error() + "无法获取mysql")
		}
		if len(table) < 1 {
			fmt.Printf("命令行错误  %s不正确\n", table) //nolint:govet
			return
		}
		gen.GenStruct(dbName, table)
		for _, t := range table {
			genDao(dbName, t, t)
		}
	},
}

func GetTables(dbName string) {
	db, _ := mysql.NewConnect(dbName)
	rows, _ := db.Raw("show tables").Rows()
	defer rows.Close()
	var tables = make([]string, 0)
	for rows.Next() {
		var name string
		rows.Scan(&name) //nolint:errcheck
		if len(strings.Split(name, "_")) < 2 {
			panic("请采用下划线模式命名数据库表名")
		}
		tables = append(tables, name)
		genDao(dbName, name, name)
	}
	gen.GenStruct(dbName, tables)
	return
}
