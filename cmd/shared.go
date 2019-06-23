package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	Certificates     []string
	Tokens           []string
	SimpleAuth       []string
	Interval         int
	ManifestFile     string
	PatchOnly        bool
	Truncate         bool
	emptyStringArray = []string{}
)

func initCredentialFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayVar(
		&SimpleAuth, "simple-auth", emptyStringArray,
		"list of credentials (user & password) prefixed by source. Ex: --simple-auth=vault01=user:password",
	)

	cmd.Flags().StringArrayVar(
		&Tokens, "tokens", emptyStringArray,
		"list of tokens prefixed by source. Ex: --tokens=vault01=wijq80jw80djqw8",
	)

	cmd.Flags().StringArrayVar(
		&Certificates, "certificates", emptyStringArray,
		"list of certificates prefixed by source. Ex: --certificates=vault01=ca://etc/certs/service,ca+crt://etc/certs/client.crt,key://etc/certs/client.pem",
	)
}

func initManifestFile(f *pflag.FlagSet) {
	f.StringVar(&ManifestFile, "manifest", "", "manifest file that dictates which files to mount to")
}

func initIOConfs(f *pflag.FlagSet) {
	f.BoolVar(&PatchOnly, "patch-only", false, "only write the files back if the value changes")
	f.BoolVar(&Truncate, "truncate", false, "use truncate to before overwriting the files")
}
