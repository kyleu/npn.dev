package parseutil

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"emperror.dev/errors"
)

func GetMatchingFiles(root string, pattern string) ([]string, error) {
	d := path.Join(root, pattern)
	if strings.Contains(pattern, "..") {
		return nil, errors.New("invalid pattern [" + pattern + "]")
	}
	return filepath.Glob(d)
}

func ReadFirstK(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	bytes := make([]byte, 1024)
	count, err := f.Read(bytes)
	if err != nil {
		return "", err
	}
	return string(bytes[0:count]), nil
}
