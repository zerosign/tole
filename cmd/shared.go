package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/zerosign/tole/base"
)

var (
	Certificates base.StringList
	Tokens       base.StringList
	SimpleAuth   base.StringList
	Interval     int
	ManifestFile string
	PatchOnly    bool
	Truncate     bool
)

func initCredentialFlags(cmd *cobra.Command) {
	cmd.Flags().Var(&SimpleAuth, "simple-auth", "list of credentials (user & password) prefixed by source. Ex: --simple-auth=vault01=user:password")
	cmd.Flags().Var(&Tokens, "tokens", "list of tokens prefixed by source. Ex: --tokens=vault01=wijq80jw80djqw8")
	cmd.Flags().Var(&Certificates, "certificates", "list of certificates prefixed by source. Ex: --certificates=vault01=ca://etc/certs/service,ca+crt://etc/certs/client.crt,key://etc/certs/client.pem")
}

func initManifestFile(f *pflag.FlagSet) {
	f.StringVar(&ManifestFile, "manifest", "", "manifest file that dictates which files to mount to")
}

func initIOConfs(f *pflag.FlagSet) {
	f.BoolVar(&PatchOnly, "patch-only", false, "only write the files back if the value changes")
	f.BoolVar(&Truncate, "truncate", false, "use truncate to before overwriting the files")
}
