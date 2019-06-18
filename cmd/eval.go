package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	evalCmd = &cobra.Command{
		Use:     "eval",
		Aliases: []string{},
		Short:   "eval a template configuration once",
		Long:    `Eval template configuration once`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("args: %v", args)
		},
	}
)

func init() {
	initCredentialFlags(evalCmd)
	initManifestFile(evalCmd.Flags())
	initIOConfs(evalCmd.Flags())
}
