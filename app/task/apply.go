package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/output"
)

type FileResult struct {
	File    *output.File
	Message string
	Written bool
	Error   error
}
type FileResults []*FileResult

func (fr FileResults) WrittenCount() int {
	y := 0
	for _, f := range fr {
		if f.Written { y++ }
	}
	return y
}

func (fr FileResults) SkippedCount() int {
	n := 0
	for _, f := range fr {
		if !f.Written { n++ }
	}
	return n
}

func (r Result) applyOutput(files []*output.File, dest string) (FileResults, error) {
	var ret FileResults

	f, err := os.Open(dest)
	if err != nil || f == nil {
		return nil, errors.Wrap(err, "destination ["+dest+"] does not exist")
	}

	for _, file := range files {
		dir := path.Join(dest, file.Pkg.ToPath())
		lines := file.Render()
		data := strings.Join(lines, "\n")

		fn := path.Join(dir, file.Filename)

		currFile, err := os.Open(fn)
		if err != nil {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return nil, errors.Wrap(err, "unable to create directories for ["+fn+"]")
			}
		}
		curr := ""
		if currFile != nil {
			b, err := ioutil.ReadFile(fn)
			if err != nil {
				return nil, errors.Wrap(err, "unable to read original file ["+fn+"]")
			}
			curr = string(b)
		}

		if len(data) == len(curr) && data == curr {
			ret = append(ret, &FileResult{File: file, Message: "skipped [" + fn + "]"})
		} else {
			err = ioutil.WriteFile(fn, []byte(data), 0644)
			if err != nil {
				return nil, errors.Wrap(err, "unable to write output to ["+fn+"]")
			}

			msg := fmt.Sprintf("wrote [%v/%v] bytes to [%v]", len(data), len(curr), fn)
			ret = append(ret, &FileResult{File: file, Message: msg, Written: true})
		}
	}

	return ret, nil
}
