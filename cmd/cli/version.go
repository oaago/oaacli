package cli

import (
	"fmt"
	"strings"

	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "v",
	Short: "oaago version 更新时间/更新版本1",
	Run: func(cmd *cobra.Command, args []string) {
		v := string(utils.RunCmd("git describe --abbrev=0 --tags", true))
		if strings.Contains(v, "v") {
			fmt.Println(strings.Replace(v, "exit status 128", "", -1))
		} else {
			fmt.Println("暂无版本")
		}
	},
}
