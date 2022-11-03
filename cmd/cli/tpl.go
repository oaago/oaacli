package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var TplCli = &cobra.Command{
	Use:     "setTpl",
	Aliases: []string{"i"},
	Short:   "oaago setTpl 根据 oaago.json 生成出来需要的项目文件， 可以制定配置文件oaago.json 别名 i 例如 oaa i",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(`
			0: "api",
			1: "service",
			2: "service-handler",
			3: "http-router",
			4: "dao",
		`)
			panic("请指定tpl 的类型和 文件模板")
		}
		run := false
		tplMap := map[int]string{
			0: "api",
			1: "service",
			2: "service-handler",
			3: "http-router",
			4: "dao",
		}
		for _, s := range tplMap {
			if s == args[0] {
				run = true
				break
			}
		}
		if run {
			SetTpl()
		}
	},
}

func SetTpl() {

}
