package npncore

import "strings"

// Handles filesystem operations (multiple implementations)
type FileLoader interface {
	// Root directory, as a string
	Root() string
	// Reads the contents of a file as a byte array
	ReadFile(path string) ([]byte, error)
	// Creates a directory, like it says on the tin
	CreateDirectory(path string) error
	// Writes the the provided byte array to a file
	WriteFile(path string, content []byte, overwrite bool) error
	// Copies the contents of one file to another
	CopyFile(src string, tgt string) error
	// Lists all files in a directory with a `.json` extension
	ListJSON(path string) []string
	// Lists all files in a directory with a provided extension
	ListExtension(path string, ext string) []string
	// Lists all directories in a directory
	ListDirectories(path string) []string
	// Returns a boolean indicating if the file exists, and another boolean to indicate if it's a directory
	Exists(path string) (bool, bool)
	// Removes the file at the provided path
	Remove(path string) error
	// Removes the file at the provided path, recursively
	RemoveRecursive(pt string) error
}

// Returns the filename from a given path
func FilenameOf(fn string) string {
	idx := strings.LastIndex(fn, "/")
	if idx > -1 {
		fn = fn[idx+1:]
	}
	return fn
}
