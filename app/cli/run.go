package cli

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnservice/auth"
	"os"
	"strings"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"emperror.dev/handler/logur"
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controllers"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	log "logur.dev/logur"
)

const Version = "0.0.1"

var FileLoaderOverride npncore.FileLoader

func init() {
	npncore.AppKey = "npn"
	npncore.AppName = npncore.AppKey
	npncore.AppVersion = Version

	auth.RequiredAuthProviders = []string{"google"}
	auth.RequiredAuthDomains = []string{"kyleu.com"}
	npncontroller.FileBrowseRoot = "./data"
}

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

	return info, r, nil
}

func InitApp(platform string, dir string) npnweb.AppInfo {
	_ = os.Setenv("TZ", "UTC")
	npncore.AppPlatform = platform

	setIcon()

	logger := npncore.InitLogging(verbose)
	logger = log.WithFields(logger, map[string]interface{}{"debug": verbose, "version": Version})

	errorHandler := logur.New(logger)
	defer emperror.HandleRecover(errorHandler)

	dir = strings.TrimSpace(dir)
	if len(dir) == 0 {
		dir = defaultDirectory()
	}

	var files npncore.FileLoader
	if FileLoaderOverride != nil {
		files = FileLoaderOverride
	} else {
		if platform == "wasm" {
			logger.Warn("can't load FileLoaderOverride for WASM")
			files = npncore.NewFileSystem(dir, logger)
		} else {
			files = npncore.NewFileSystem(dir, logger)
		}
	}

	return app.NewService(verbose, files, redir, logger)
}

func setIcon() {
	npnweb.IconContent = `<svg width="32px" height="32px" viewBox="-0 0 68 68" xmlns="http://www.w3.org/2000/svg">
            <g fill="none">
              <path id="logo-symbol" d="M 50.655 0 L 50.611 12.31 L 30.603 26.31 M 30.603 42.31 L 50.611 56.31 L 50.611 68.048 M 2 34.371 L 28.902 34.31 M 28.902 17.31 L 28.902 51.31 M 30.303 51.31 L 30.303 17.31 M 40.607 52.31 L 43.609 48.31 L 47.61 54.31 L 40.607 52.31 Z M 8.594 33.31 C 9.364 11.769 33.173 -0.86 51.451 10.577 C 69.728 22.014 68.766 48.94 49.718 59.043 C 31.449 68.734 9.332 55.971 8.594 35.31" style="stroke-width: 5px; paint-order: fill; stroke: rgb(135, 135, 135);" />
            </g>
          </svg>`
}
