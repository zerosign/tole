package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	watchCmd = &cobra.Command{
		Use:     "watch",
		Aliases: []string{},
		Short:   "watch a template configuration changes",
		Long: `watch sources and apply any changes related to
its variables into output based on templates`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("args: %v", args)
		},
	}
)

func init() {
	initCredentialFlags(watchCmd)
	initManifestFile(watchCmd.Flags())
	initIOConfs(watchCmd.Flags())
	watchCmd.Flags().IntVar(
		&Interval, "interval", 15,
		`watch interval in seconds, if the backend support
watch it will still try to sync itself gradually even
there is no event from the storage backend`,
	)
}
