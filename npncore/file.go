package npncore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"emperror.dev/errors"
	"logur.dev/logur"
)

type FileLoader struct {
	root   string
	logger logur.Logger
}

func FilenameOf(fn string) string {
	idx := strings.LastIndex(fn, "/")
	if idx > -1 {
		fn = fn[idx+1:]
	}
	return fn
}

func NewFileLoader(root string, logger logur.Logger) *FileLoader {
	return &FileLoader{root: root, logger: logger}
}

func (f *FileLoader) getPath(ss ...string) string {
	s := path.Join(ss...)
	if strings.HasPrefix(s, f.root) {
		return s
	}
	return path.Join(f.root, s)
}

func (f *FileLoader) Root() string {
	return f.root
}

func (f *FileLoader) ReadFile(path string) (string, error) {
	b, err := ioutil.ReadFile(f.getPath(path))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (f *FileLoader) WriteFile(path string, content string, overwrite bool) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) && !overwrite {
		return errors.New("file exists, will not overwrite")
	}
	dd := filepath.Dir(f.getPath(path))
	err = os.MkdirAll(dd, 0755)
	if err != nil {
		return errors.Wrap(err, "unable to create data directory ["+dd+"]")
	}
	file, err := os.Create(f.getPath(path))
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(content))
	return err
}

func (f *FileLoader) ListJSON(path string) []string {
	matches, err := filepath.Glob(f.getPath(path, "*.json"))
	if err != nil {
		f.logger.Warn(fmt.Sprintf("cannot list JSON in path ["+path+"]: %+v", err))
	}
	ret := make([]string, 0, len(matches))
	for _, j := range matches {
		idx := strings.LastIndex(j, "/")
		if idx > 0 {
			j = j[idx+1:]
		}
		ret = append(ret, strings.TrimSuffix(j, ".json"))
	}
	return ret
}

func (f *FileLoader) ListDirectories(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		f.logger.Warn(fmt.Sprintf("cannot list path ["+path+"]: %+v", err))
	}
	ret := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			ret = append(ret, f.Name())
		}
	}
	return ret
}

func (f *FileLoader) Remove(path string) error {
	p := f.getPath(path)
	f.logger.Warn("removing file at path [" + p + "]")
	return os.Remove(p)
}
