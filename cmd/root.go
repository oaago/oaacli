package cmd

import (
	"github.com/oaago/oaago/cmd/cli"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:     "oaago",
		Aliases: []string{"oaa"},
		Short:   "A generator for oaago 别名:  oaa",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cli.NewProject)
	rootCmd.AddCommand(cli.VersionCmd)
	rootCmd.AddCommand(cli.GenInit)
	rootCmd.AddCommand(cli.UpdateCmd)
	rootCmd.AddCommand(cli.GenTable)
	rootCmd.AddCommand(cli.GenClean)
	rootCmd.AddCommand(cli.UpdateAllCmd)
	rootCmd.AddCommand(cli.DockerFileCmd)
	rootCmd.AddCommand(cli.DockerBuildCmd)
}
