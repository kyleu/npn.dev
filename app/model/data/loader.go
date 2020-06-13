package data

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"io/ioutil"
	"logur.dev/logur"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileLoader struct {
	root   string
	logger logur.Logger
}

func NewFileLoader(logger logur.Logger) *FileLoader {
	return &FileLoader{root: "." + util.AppName, logger: logger}
}

func (f *FileLoader) LoadProfile() (*util.UserProfile, error) {
	content, err := f.readFile("profile.json")
	if err != nil {
		return util.NewUserProfile(), nil
	}
	tgt := &util.UserProfile{}
	util.FromJSON([]byte(content), tgt, f.logger)
	return tgt, nil
}

func (f *FileLoader) SaveProfile(p *util.UserProfile) error {
	return f.writeFile("profile.json", util.ToJSON(p, f.logger), true)
}

func (f *FileLoader) ListSchemata() []string {
	return f.listJSON("schema")
}

func (f *FileLoader) LoadSchema(key string) (*schema.Schema, error) {
	content, err := f.readFile("schema/" + key + ".json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to find schema file with key [" + key + "]")
	}
	tgt := &schema.Schema{}
	util.FromJSON([]byte(content), tgt, f.logger)
	return tgt, nil
}

func (f *FileLoader) SaveSchema(sch *schema.Schema, overwrite bool) error {
	return f.writeFile("schema/"+sch.Key+".json", util.ToJSON(sch, f.logger), overwrite)
}

func (f *FileLoader) getPath(ss ...string) string {
	s := path.Join(ss...)
	if strings.HasPrefix(s, f.root) {
		return s
	}
	return path.Join(f.root, s)
}

func (f *FileLoader) readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(f.getPath(path))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (f *FileLoader) writeFile(path string, content string, overwrite bool) error {
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

func (f *FileLoader) listJSON(path string) []string {
	matches, err := filepath.Glob(f.getPath(path, "*.json"))
	if err != nil {
		f.logger.Warn(fmt.Sprintf("cannot list JSON in path [" + path + "]: %+v", err))
	}
	ret := make([]string, 0, len(matches))
	for _, j := range matches {
		idx := strings.LastIndex(j, "/")
		if idx > 0 {
			j = j [idx+1:]
		}
		ret = append(ret, strings.TrimSuffix(j, ".json"))
	}
	return ret
}
