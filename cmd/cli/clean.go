package cli

import (
	"github.com/spf13/cobra"
	"os"
)

var GenClean = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "oaacli init 根据 oaago.json 生成出来需要的项目文件， 可以制定配置文件oaago.json 别名 c 例如 oaa c",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		os.RemoveAll("./internal/router/")
		os.RemoveAll("./internal/api/http/")
		os.RemoveAll("./internal/api/rpc/")
		os.RemoveAll("./internal/dao/")
	},
}
