package cli

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controllers"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	log "logur.dev/logur"
	"os"
	"strings"
)

func Run(a string, p uint16, dir string, version string, commitHash string) (uint16, error) {
	info, r, err := Start(dir, version, commitHash)
	if err != nil {
		return p, err
	}

	return npnweb.MakeServer(info, r, a, p)
}

func Start(dir string, version string, commitHash string) (npnweb.AppInfo, *mux.Router, error) {
	info := InitApp(dir, version, commitHash)

	r, err := controllers.BuildRouter(info)
	if err != nil {
		return info, nil, errors.WithMessage(err, "unable to construct routes")
	}

	setContent()

	return info, r, nil
}

func InitApp(dir string, version string, commitHash string) npnweb.AppInfo {
	_ = os.Setenv("TZ", "UTC")

	npncore.AppKey = "npn"
	npncore.AppName = npncore.AppKey

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": version, "commit": commitHash})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	dir = strings.TrimSpace(dir)
	if len(dir) == 0 {
		dir = defaultDirectory()
	}

	return app.NewService(verbose, dir, version, commitHash, logger)
}
