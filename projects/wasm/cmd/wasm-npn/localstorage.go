package main

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"path"
	"strings"
	"syscall/js"
)

type LocalStorageLoader struct {
	storage js.Value
}

func NewLocalStorageLoader() npncore.FileLoader {
	return &LocalStorageLoader{storage: js.Global().Get("localStorage")}
}

func (l *LocalStorageLoader) getPath(ss ...string) string {
	s := path.Join(ss...)
	return path.Join(l.Root(), s)
}

func (l *LocalStorageLoader) Root() string {
	return npncore.AppKey
}

func (l *LocalStorageLoader) ReadFile(path string) ([]byte, error) {
	path = l.getPath(path)
	res := l.storage.Call("getItem", path)
	if res.IsNull() || res.IsUndefined() {
		return nil, errors.New("file [" + path + "] not found")
	}
	return []byte(res.String()), nil
}

func (l *LocalStorageLoader) CreateDirectory(path string) error {
	return nil
}

func (l *LocalStorageLoader) WriteFile(path string, content []byte, overwrite bool) error {
	path = l.getPath(path)
	l.storage.Call("setItem", path, string(content))
	return nil
}

func (l *LocalStorageLoader) CopyFile(src string, tgt string) error {
	sp := l.getPath(src)
	tp := l.getPath(tgt)

	targetExists, _ := l.Exists(tp)
	if targetExists {
		return errors.New("file [" + tp + "] exists, will not overwrite")
	}

	input, err := l.ReadFile(sp)
	if err != nil {
		return err
	}

	err = l.WriteFile(tp, input, false)
	return err
}

func (l *LocalStorageLoader) ListJSON(path string) []string {
	return l.ListExtension(path, "json")
}

func (l *LocalStorageLoader) ListExtension(path string, ext string) []string {
	path = l.getPath(path)
	ret := make([]string, 0)
	for _, key := range l.Keys() {
		if(strings.HasPrefix(key, path)) {
			tgt := strings.TrimPrefix(key, path)
			tgt = strings.TrimPrefix(tgt, "/")
			if !strings.Contains(tgt, "/") {
				if strings.HasSuffix(tgt, ext) {
					exists := false
					for _, x := range ret {
						if x == tgt {
							exists = true
						}
					}
					if !exists {
						ret = append(ret, tgt)
					}
				}
			}
		}
	}
	return ret
}

func (l *LocalStorageLoader) ListDirectories(path string) []string {
	path = l.getPath(path)
	ret := make([]string, 0)
	for _, key := range l.Keys() {
		if(strings.HasPrefix(key, path)) {
			tgt := strings.TrimPrefix(key, path)
			tgt = strings.TrimPrefix(tgt, "/")
			tgt, _ = npncore.SplitString(tgt, '/', true)
			exists := false
			for _, x := range ret {
				if x == tgt {
					exists = true
				}
			}
			if !exists {
				ret = append(ret, tgt)
			}
		}
	}
	return ret
}

func (l *LocalStorageLoader) Exists(path string) (bool, bool) {
	path = l.getPath(path)
	matched := false
	exact := false
	for _, key := range l.Keys() {
		if(strings.HasPrefix(key, path)) {
			matched = true
			if key == path {
				exact = true
			}
		}
	}
	isDir := matched && (!exact)
	// log(fmt.Sprintf("Exists(%v): %v, IsDir: %v", path, matched, isDir))
	return matched, isDir
}

func (l *LocalStorageLoader) Remove(path string) error {
	l.storage.Call("removeItem", path)
	return nil
}

func (l *LocalStorageLoader) RemoveRecursive(pt string) error {
	return errors.New(fmt.Sprintf("TODO: RemoveRecursive(%v)", pt))
}

func (l *LocalStorageLoader) Keys() []string {
	ret := make([]string, 0)
	for i := 0; i < l.storage.Length(); i++ {
		key := l.storage.Call("key", i).String()
		ret = append(ret, key)
	}
	return ret
}
