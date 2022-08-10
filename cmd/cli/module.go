package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ModuleCmd = &cobra.Command{
	Use:   "gm",
	Short: "获取 mod module 名称",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(module)
	},
}
