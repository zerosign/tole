package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zerosign/tole/manifest"
	"github.com/zerosign/tole/source"
	"log"
	"os"
)

var (
	BackendCheck bool

	checkCmd = &cobra.Command{
		Use:     "check",
		Aliases: []string{},
		Short:   "check for source readiness",
		Run: func(cmd *cobra.Command, args []string) {
			manifest, errors := manifest.ParseManifest(ManifestFile)

			if len(errors) > 0 {
				log.Fatal("manifest error detected, please run lint to check which wrong")
			}

			sources, errors := source.Registry(manifest.Sources())

			if len(errors) > 0 {
				fmt.Errorf("some sources can't be contacted.\n")
				for idx, err := range errors {
					fmt.Errorf("%d: %s\n", idx, err)
				}
				os.Exit(1)
			}

			if sources.IsEmpty() {
				log.Fatal("no sources detected.")
			} else {
				// assert(len(sources) == len(manifest.Sources()))
				if sources.Size() == len(manifest.Sources()) {
					log.Fatal("sources declared in manifest should be equals to sources factories")
				}
				log.Print("no errors detected, most sources are contactable")
			}

		},
	}
)

func init() {
	initManifestFile(checkCmd.Flags())
	checkCmd.Flags().BoolVar(&BackendCheck, "--backend", true, "check backend availabitity")
}

// check whether the backend exists or not
// check whether we can do auth to the backend or not
// check whether template sources are readable/exists or not
// check whether target mounts are writeable or not
