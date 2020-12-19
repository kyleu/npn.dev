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

// Implements `FileLoader` for local filesystem access
type FileSystem struct {
	root   string
	logger logur.Logger
}

var _ FileLoader = (*FileSystem)(nil)

// Constructor
func NewFileSystem(root string, logger logur.Logger) *FileSystem {
	return &FileSystem{root: root, logger: logger}
}

func (f *FileSystem) getPath(ss ...string) string {
	s := path.Join(ss...)
	if strings.HasPrefix(s, f.root) {
		return s
	}
	return path.Join(f.root, s)
}

// Root directory, as a string
func (f *FileSystem) Root() string {
	return f.root
}

// Reads the contents of a file as a byte array
func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(f.getPath(path))
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Creates a directory, like it says on the tin
func (f *FileSystem) CreateDirectory(path string) error {
	p := f.getPath(path)
	err := os.MkdirAll(p, 0755)
	if err != nil {
		return errors.Wrap(err, "unable to create data directory ["+p+"]")
	}
	return nil
}

// Writes the the provided byte array to a file
func (f *FileSystem) WriteFile(path string, content []byte, overwrite bool) error {
	p := f.getPath(path)
	_, err := os.Stat(p)
	if os.IsExist(err) && !overwrite {
		return errors.New("file [" + p + "] exists, will not overwrite")
	}
	dd := filepath.Dir(path)
	err = f.CreateDirectory(dd)
	if err != nil {
		return errors.Wrap(err, "unable to create data directory ["+dd+"]")
	}
	file, err := os.Create(p)
	if err != nil {
		return errors.Wrap(err, "unable to create file ["+p+"]")
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(content)
	if err != nil {
		return errors.Wrap(err, "unable to write content to file ["+p+"]")
	}
	return nil
}

// Copies the contents of one file to another
func (f *FileSystem) CopyFile(src string, tgt string) error {
	sp := f.getPath(src)
	tp := f.getPath(tgt)

	targetExists, _ := f.Exists(tp)
	if targetExists {
		return errors.New("file [" + tp + "] exists, will not overwrite")
	}

	input, err := f.ReadFile(sp)
	if err != nil {
		return err
	}

	err = f.WriteFile(tp, input, false)
	return err
}

// Lists all files in a directory with a `.json` extension
func (f *FileSystem) ListJSON(path string) []string {
	return f.ListExtension(path, "json")
}

// Lists all files in a directory with a provided extension
func (f *FileSystem) ListExtension(path string, ext string) []string {
	glob := "*." + ext
	matches, err := filepath.Glob(f.getPath(path, glob))
	if err != nil {
		f.logger.Warn(fmt.Sprintf("cannot list ["+ext+"] in path ["+path+"]: %+v", err))
	}
	ret := make([]string, 0, len(matches))
	for _, j := range matches {
		idx := strings.LastIndex(j, "/")
		if idx > 0 {
			j = j[idx+1:]
		}
		ret = append(ret, strings.TrimSuffix(j, "."+ext))
	}
	return ret
}

// Lists all directories in a directory
func (f *FileSystem) ListDirectories(path string) []string {
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

// Returns a boolean indicating if the file exists, and another boolean to indicate if it's a directory
func (f *FileSystem) Exists(path string) (bool, bool) {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err == nil {
		return true, s.IsDir()
	}
	return false, false
}

// Removes the file at the provided path
func (f *FileSystem) Remove(path string) error {
	p := f.getPath(path)
	f.logger.Warn("removing file at path [" + p + "]")
	return os.Remove(p)
}

// Removes the file at the provided path, recursively
func (f *FileSystem) RemoveRecursive(pt string) error {
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
