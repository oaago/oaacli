package cli

import (
	"fmt"

	"github.com/oaago/oaago/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update oaago version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("将会自动更新到master分支 请稍等....")
		utils.RunCmd("go install github.com/oaago/oaago@mian", true)
		utils.RunCmd("go install github.com/oaago/protoc-gen-oaago@mian", true)
		fmt.Println("更新完成")
	},
}

var UpdateAllCmd = &cobra.Command{
	Use:   "updateall",
	Short: "update oaago version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("将会全部自动更新到main分支 请稍等....")
		utils.RunCmd("go install github.com/oaago/oaago@mian", true)
		utils.RunCmd("go install github.com/oaago/protoc-gen-oaago@mian", true)
		utils.RunCmd("go install github.com/oaago/server@mian", true)
		utils.RunCmd("go install github.com/oaago/cloud@mian", true)
		utils.RunCmd("go install github.com/oaago/common@mian", true)
		fmt.Println("更新完成")
	},
}
