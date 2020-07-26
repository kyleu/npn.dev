package cli

import (
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncore"
	"github.com/spf13/cobra"
)

var verbose bool
var redir string
var addr string
var port uint16

// Configure configures a root command.
func Configure(version string, commitHash string) cobra.Command {
	rootCmd := cobra.Command{
		Use:   npncore.AppName,
		Short: "Command line interface for " + npncore.AppName,
		RunE: func(cmd *cobra.Command, _ []string) error {
			info, err := InitApp(version, commitHash)
			if err != nil {
				return errors.Wrap(err, "error initializing application")
			}

			return MakeServer(info, addr, port)
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&redir, "redir", "r", "http://localhost:10101", "redirect url for signin, defaults to localhost")
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 10101, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	return rootCmd
}
