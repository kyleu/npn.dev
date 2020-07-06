package web

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
	"logur.dev/logur"

	"github.com/kyleu/npn/app/util"

	"github.com/gorilla/sessions"
)

type Breadcrumb struct {
	Path  string
	Title string
}

func BreadcrumbSelf(title string) Breadcrumb {
	return Breadcrumb{Path: "", Title: title}
}

type Breadcrumbs []Breadcrumb

func BreadcrumbsSimple(path string, title string) Breadcrumbs {
	return []Breadcrumb{
		{path, title},
	}
}

var sessionKey = func() string {
	x := os.Getenv("SESSION_KEY")
	if len(x) == 0 {
		x = "random_secret_key"
	}
	return x
}()

var store = sessions.NewCookieStore([]byte(sessionKey))

const sessionName = util.AppName + "-session"

func ParseFlash(s string) (string, string) {
	split := strings.SplitN(s, ":", 2)
	severity := split[0]
	content := split[1]

	switch severity {
	case util.KeyStatus:
		return "uk-alert-primary", content
	case "success":
		return "uk-alert-success", content
	case "warning":
		return "uk-alert-warning", content
	case util.KeyError:
		return "uk-alert-danger", content
	default:
		return "", content
	}
}

var re *regexp.Regexp

func PathParams(s string) []string {
	if re == nil {
		re = regexp.MustCompile("{([^}]*)}")
	}

	matches := re.FindAll([]byte(s), -1)

	ret := make([]string, 0, len(matches))
	for _, m := range matches {
		ret = append(ret, string(m))
	}

	return ret
}

func Route(routes *mux.Router, logger logur.Logger, act string, pairs ...string) string {
	route := routes.Get(act)
	if route == nil {
		msg := "cannot find route at path [" + act + "]"
		logger.Warn(fmt.Sprintf("%v: %+v", msg, errors.New(msg)))
		return "/route/notfound/" + act
	}
	u, err := route.URL(pairs...)
	if err != nil {
		msg := "cannot bind route at path [" + act + "]"
		logger.Warn(fmt.Sprintf("%v: %+v", msg, errors.New(msg)))
		return "/route/error/" + act
	}
	return u.Path
}
