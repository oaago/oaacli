package cmd

import (
	"github.com/oaago/oaacli/cmd/cli"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:     "oaacli",
		Aliases: []string{"oaa"},
		Short:   "A generator for oaacli 别名:  oaa",
	}
)

func Execute() {
	rootCmd.AddCommand(cli.NewProject)
	rootCmd.AddCommand(cli.VersionCmd)
	//rootCmd.AddCommand(cli.GenApi)
	rootCmd.AddCommand(cli.GenInit)
	//rootCmd.AddCommand(cli.GenType)

	//rootCmd.AddCommand(cli.GenDao)
	rootCmd.AddCommand(cli.GenService)
	//rootCmd.AddCommand(cli.GenModel)
	rootCmd.AddCommand(cli.UpdateCmd)
	//rootCmd.AddCommand(cli.ModuleCmd)
	rootCmd.AddCommand(cli.GenTable)
	rootCmd.AddCommand(cli.GenClean)
	rootCmd.AddCommand(cli.GenRpc)
	//rootCmd.AddCommand(cli.GenRpcAdd)
}
