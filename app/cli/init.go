package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/schema"

	"github.com/kyleu/npn/app/parser"

	"github.com/kyleu/npn/app/web/routes"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/gorilla/handlers"
	"github.com/kyleu/npn/app/config"
	"github.com/kyleu/npn/app/util"
	log "logur.dev/logur"
)

func InitApp(version string, commitHash string) (*config.AppInfo, error) {
	_ = os.Setenv("TZ", "UTC")

	logger := initLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	ai := initAppInfo(logger, version, commitHash)

	return ai, nil
}

func initAppInfo(logger log.Logger, version string, commitHash string) *config.AppInfo {
	return &config.AppInfo{
		Debug:    verbose,
		Parsers:  parser.NewParsers(logger),
		Files:    util.NewFileLoader(logger),
		Schemata: schema.NewCache(logger),
		Projects: project.NewCache(logger),
		Version:  version,
		Commit:   commitHash,
		Logger:   logger,
	}
}

func MakeServer(info *config.AppInfo, address string, port uint16) error {
	r, err := routes.BuildRouter(info)
	if err != nil {
		return errors.WithMessage(err, "unable to construct routes")
	}
	var msg = "%v is starting on [%v:%v]"
	if info.Debug {
		msg += " (verbose)"
	}
	info.Logger.Info(fmt.Sprintf(msg, util.AppName, address, port))
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), handlers.CORS()(r))
	return errors.Wrap(err, "unable to run http server")
}
