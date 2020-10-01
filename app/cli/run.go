package cli

import (
	"os"
	"strings"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controllers"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	log "logur.dev/logur"
)

const Version = "0.0.1"

var FileLoaderOverride npncore.FileLoader

func Run(a string, p uint16, platform string, dir string) (uint16, error) {
	info, r, err := Start(platform, dir)
	if err != nil {
		return p, err
	}

	return npnweb.MakeServer(info, r, a, p)
}

func Start(platform string, dir string) (npnweb.AppInfo, *mux.Router, error) {
	info := InitApp(platform, dir)

	r, err := controllers.BuildRouter(info)
	if err != nil {
		return info, nil, errors.WithMessage(err, "unable to construct routes")
	}

	setContent()

	return info, r, nil
}

func InitApp(platform string, dir string) npnweb.AppInfo {
	_ = os.Setenv("TZ", "UTC")

	npncore.AppKey = "npn"
	npncore.AppName = npncore.AppKey
	npncore.AppPlatform = platform
	npncore.AppVersion = Version

	npncontroller.FileBrowseRoot = "./data"

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": Version})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	dir = strings.TrimSpace(dir)
	if len(dir) == 0 {
		dir = defaultDirectory()
	}

	var files npncore.FileLoader
	if platform == "wasm" {
		if FileLoaderOverride == nil {
			logger.Warn("can't load FileLoaderOverride for WASM")
			files = npncore.NewFileSystem(dir, logger)
		} else {
			files = FileLoaderOverride
		}
	} else {
		files = npncore.NewFileSystem(dir, logger)
	}

	return app.NewService(verbose, files, logger)
}
