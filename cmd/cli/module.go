package cli

import (
	"fmt"
	"strings"

	"github.com/oaago/oaacli/utils"
	"github.com/spf13/cobra"
)

var ModuleCmd = &cobra.Command{
	Use:   "gm",
	Short: "获取 mod module 名称",
	Run: func(cmd *cobra.Command, args []string) {
		v := string(utils.RunCmd("go list -m", true))
		if strings.Contains(v, "v") {
			fmt.Println(v)
		}
	},
}
