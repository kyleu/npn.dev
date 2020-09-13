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
	p := f.getPath(path)
	dd := filepath.Dir(p)
	err = os.MkdirAll(dd, 0755)
	if err != nil {
		return errors.Wrap(err, "unable to create data directory ["+dd+"]")
	}
	file, err := os.Create(p)
	if err != nil {
		return errors.Wrap(err, "unable to create file ["+p+"]")
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write([]byte(content))
	if err != nil {
		return errors.Wrap(err, "unable to write content to file ["+p+"]")
	}
	if !strings.HasSuffix(content, "\n") {
		_, err = file.Write([]byte("\n"))
		if err != nil {
			return errors.Wrap(err, "unable to write ending linebreak to file ["+p+"]")
		}
	}
	return nil
}

func (f *FileLoader) ListJSON(path string) []string {
	return f.ListExtension(path, "json")
}

func (f *FileLoader) ListExtension(path string, ext string) []string {
	glob := "*." + ext
	matches, err := filepath.Glob(f.getPath(path, glob))
	if err != nil {
		f.logger.Warn(fmt.Sprintf("cannot list [" + ext + "] in path ["+path+"]: %+v", err))
	}
	ret := make([]string, 0, len(matches))
	for _, j := range matches {
		idx := strings.LastIndex(j, "/")
		if idx > 0 {
			j = j[idx+1:]
		}
		ret = append(ret, strings.TrimSuffix(j, "." + ext))
	}
	return ret
}

func (f *FileLoader) ListDirectories(path string) []string {
	p := f.getPath(path)
	files, err := ioutil.ReadDir(p)
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

func (f *FileLoader) Exists(path string) (bool, bool) {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err == nil {
		return true, s.IsDir()
	}
	return false, false
}

func (f *FileLoader) Remove(path string) error {
	p := f.getPath(path)
	f.logger.Warn("removing file at path [" + p + "]")
	return os.Remove(p)
}

func (f *FileLoader) RemoveRecursive(pt string) error {
	p := f.getPath(pt)
	s, err := os.Stat(p)
	if err != nil {
		return err
	}
	if s.IsDir() {
		files, err := ioutil.ReadDir(p)
		if err != nil {
			f.logger.Warn(fmt.Sprintf("cannot list path ["+pt+"]: %+v", err))
		}
		for _, file := range files {
			err := f.RemoveRecursive(path.Join(pt, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return os.Remove(p)
}
