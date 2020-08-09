package main

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controllers"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/spf13/cobra"
	log "logur.dev/logur"
	"os"
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
			npncore.AppKey = "npn"
			npncore.AppName = "npn"

			info, err := initApp(version, commitHash)
			if err != nil {
				return errors.Wrap(err, "error initializing application")
			}

			r, err := controllers.BuildRouter(info)
			if err != nil {
				return errors.WithMessage(err, "unable to construct routes")
			}

			setIcon()

			return npnweb.MakeServer(info, r, addr, port)
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&redir, "redir", "r", "http://localhost:10101", "redirect url for signin, defaults to localhost")
	flags.StringVarP(&addr, "address", "a", "127.0.0.1", "interface address to listen on")
	flags.Uint16VarP(&port, "port", "p", 10101, "port for http server to listen on")
	flags.BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	return rootCmd
}

func initApp(version string, commitHash string) (npnweb.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	npncore.AppName = "npn"

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	return app.NewService(verbose, version, commitHash, logger), nil
}

func setIcon() {
	npnweb.IconContent = `<svg width="32px" height="32px" viewBox="-0 0 68 68" xmlns="http://www.w3.org/2000/svg">
		<g fill="none">
			<path id="logo-symbol" d="M 50.655 0 L 50.611 12.31 L 30.603 26.31 M 30.603 42.31 L 50.611 56.31 L 50.611 68.048 M 2 34.371 L 28.902 34.31 M 28.902 17.31 L 28.902 51.31 M 30.303 51.31 L 30.303 17.31 M 40.607 52.31 L 43.609 48.31 L 47.61 54.31 L 40.607 52.31 Z M 8.594 33.31 C 9.364 11.769 33.173 -0.86 51.451 10.577 C 69.728 22.014 68.766 48.94 49.718 59.043 C 31.449 68.734 9.332 55.971 8.594 35.31" style="stroke-width: 5px; paint-order: fill; stroke: rgb(135, 135, 135);"/>
		</g>
	</svg>`
}
