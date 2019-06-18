package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zerosign/tole/manifest"
)

var (
	PrintDetails bool
	TemplateLint bool
	lintCmd      = &cobra.Command{
		Use:     "lint",
		Aliases: []string{},
		Short:   "check for manifest & template correctness",
		Run: func(cmd *cobra.Command, args []string) {
			// check for manifest correctness
			manifest, errors := manifest.ParseManifest(ManifestFile)

			if len(errors) > 0 {
				fmt.Println("errors detected in : ")

				for idx, err := range errors {
					fmt.Printf("%d: %v\n", idx, err)
				}
			} else {
				fmt.Println("no error detected.")

				if PrintDetails {
					FormatManifest(manifest)
				}
			}

			if TemplateLint {
				// check for template correctness
				for _, mount := range manifest.Mounts() {
					fmt.Printf("target: %s", mount.Target())
				}
			}
		},
	}
)

func FormatManifest(m manifest.Manifest) {
	var sources, aliases, mounts string
	// TODO: @zerosign
	fmt.Printf(`
manifest version: %s

sources:
%s

aliases:
%s

mounts:
%s
`, m.Version(), sources, aliases, mounts)
}

func init() {
	initManifestFile(lintCmd.Flags())
	lintCmd.Flags().BoolVar(&PrintDetails, "print-details", false, "print details about the manifest")
	lintCmd.Flags().BoolVar(&TemplateLint, "lint-template", false, "also lint the template")
}
