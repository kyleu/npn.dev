package cli

import (
	"github.com/kirsle/configdir"
	"github.com/kyleu/libnpn/npncore"
	"github.com/spf13/cobra"
)

var verbose bool
var public bool
var multiuser bool
var secret string
var redir string
var addr string
var port uint16
var dataDir string

func Configure() cobra.Command {
	InitKeys()

	rootCmd := cobra.Command{
		Use:   npncore.AppKey,
		Short: "Command line interface for " + npncore.AppName,
		RunE: func(cmd *cobra.Command, _ []string) error {
			actualPort, err := Run(addr, port, "", dataDir)
			if actualPort > 0 {
				port = actualPort
			}
			return err
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&dataDir, "datadir", "d", "", "directory to load configuration from")
	flags.StringVarP(&redir, "redir", "r", "http://localhost:10101", "redirect url for signin, defaults to localhost")
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 10101, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "when set, displays verbose output")
	flags.BoolVarP(&multiuser, "multiuser", "m", false, "when set, allows multiple users to access the server")
	flags.BoolVar(&public, "public", false, "if set, renders as public server")
	flags.StringVarP(&secret, "secret", "s", "", "protect the site with this super-secret password")

	return rootCmd
}

func defaultDirectory() string {
	dir := configdir.LocalConfig(npncore.AppName)
	_ = configdir.MakePath(dir) // Ensure it exists.
	return dir
}
