package npncore

import "strings"

type FileLoader interface {
	Root() string
	ReadFile(path string) ([]byte, error)
	CreateDirectory(path string) error
	WriteFile(path string, content []byte, overwrite bool) error
	CopyFile(src string, tgt string) error
	ListJSON(path string) []string
	ListExtension(path string, ext string) []string
	ListDirectories(path string) []string
	Exists(path string) (bool, bool)
	Remove(path string) error
	RemoveRecursive(pt string) error
}

func FilenameOf(fn string) string {
	idx := strings.LastIndex(fn, "/")
	if idx > -1 {
		fn = fn[idx+1:]
	}
	return fn
}
