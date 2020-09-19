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

func Run(a string, p uint16, platform string, dir string, version string, commitHash string) (uint16, error) {
	info, r, err := Start(platform, dir, version, commitHash)
	if err != nil {
		return p, err
	}

	return npnweb.MakeServer(info, r, a, p)
}

func Start(platform string, dir string, version string, commitHash string) (npnweb.AppInfo, *mux.Router, error) {
	info := InitApp(platform, dir, version, commitHash)

	r, err := controllers.BuildRouter(info)
	if err != nil {
		return info, nil, errors.WithMessage(err, "unable to construct routes")
	}

	setContent()

	return info, r, nil
}

func InitApp(platform string, dir string, version string, commitHash string) npnweb.AppInfo {
	_ = os.Setenv("TZ", "UTC")

	npncore.AppKey = "npn"
	npncore.AppName = npncore.AppKey
	npncore.AppPlatform = platform

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
