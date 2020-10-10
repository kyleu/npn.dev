package npncore

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"
)

type appFormatter struct {
	nested *logrus.TextFormatter
}

var root, _ = filepath.Abs(".")

func (a *appFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Message = "\n" + entry.Message + "\n\n"
	b, err := a.nested.Format(entry)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	var ret []string
	if len(lines) > 3 {
		header := lines[0]
		idx, file, line := 6, "", 0
		for {
			_, f, l, _ := runtime.Caller(idx)
			if strings.Contains(f, "logur") || strings.Contains(f, "logrus") || strings.Contains(f, "logging") {
				idx++
			} else {
				file, line = f, l
				file = strings.TrimPrefix(file, root+"/")
				break
			}
		}

		footer := strings.TrimSpace(lines[len(lines)-2])

		ret = append(ret, fmt.Sprintf(" :: %v%v:%v [%v]", header, file, line, footer))

		content := lines[1 : len(lines)-2]
		for _, s := range content {
			if len(strings.TrimSpace(s)) > 0 {
				ret = append(ret, s)
			}
		}

		ret = append(ret, lines[len(lines)-1])
	}
	return []byte(strings.Join(ret, "\n")), nil
}

func InitLogging(verbose bool) logur.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	if verbose {
		logger.SetFormatter(&appFormatter{nested: &logrus.TextFormatter{}})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logrusadapter.New(logger)
}
