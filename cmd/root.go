package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tole",
		Short: "tole configuration daemon",
		Long: `tole is a daemon/cli tool that eval or
manage your configuration needs in any environment
including local envs into production`,
	}
)

func init() {
	rootCmd.AddCommand(lintCmd)
	rootCmd.AddCommand(evalCmd)
	rootCmd.AddCommand(watchCmd)
}

func Execute() {
	rootCmd.Execute()
}
