package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/model/output"
)

type OutputResult struct {
	File    *output.File
	Message string
	Written bool
	Error   error
}
type OutputResults []OutputResult

func (r Result) applyOutput(files []*output.File, dest string) (OutputResults, error) {
	var ret OutputResults

	f, err := os.Open(dest)
	if err != nil || f == nil {
		return nil, errors.Wrap(err, "destination ["+dest+"] does not exist")
	}

	for _, file := range files {
		dir := path.Join(dest, file.Pkg.ToPath())
		lines := file.Render()
		data := []byte(strings.Join(lines, "\n"))

		fn := path.Join(dir, file.Filename)

		currFile, err := os.Open(fn)
		if err != nil {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return nil, errors.Wrap(err, "unable to create directories for ["+fn+"]")
			}
		}
		var curr []byte
		if currFile != nil {
			curr, err = ioutil.ReadFile(fn)
			if err != nil {
				return nil, errors.Wrap(err, "unable to read original file ["+fn+"]")
			}
		}

		if len(data) == len(curr) && string(data) == string(curr) {
			ret = append(ret, OutputResult{File: file, Message: "skipped [" + fn + "]"})
		} else {
			err = ioutil.WriteFile(fn, data, 0644)
			if err != nil {
				return nil, errors.Wrap(err, "unable to write output to ["+fn+"]")
			}

			msg := fmt.Sprintf("wrote [%v] bytes to [%v]", len(data), fn)
			ret = append(ret, OutputResult{File: file, Message: msg})
		}
	}

	msg := fmt.Sprintf("applied [%v] files to [%v]", len(files), dest)
	ret = append(ret, OutputResult{File: nil, Message: msg})

	return ret, nil
}
