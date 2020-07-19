package bootstrap

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"html/template"
	"io/ioutil"
	"logur.dev/logur"
	"os"
	"path"
	"path/filepath"

	"emperror.dev/errors"
	"github.com/kyleu/npn/app/project"
	"github.com/kyleu/npn/app/web/assets"
)

func Extract(prototype *Prototype, cfg *project.Project, logger logur.Logger) error {
	err := os.MkdirAll(cfg.RootPath, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "destination ["+cfg.RootPath+"] does not exist")
	}

	zipFilename := "/bootstrap/" + prototype.Key + ".zip"
	zippedBytes, _, _, err := assets.Asset("web/assets", zipFilename)
	if err != nil {
		return errors.Wrap(err, "can't read asset ["+zipFilename+"]")
	}

	rx, err := gzip.NewReader(bytes.NewReader(zippedBytes))
	if err != nil {
		return err
	}
	result, _ := ioutil.ReadAll(rx)

	r, err := zip.NewReader(bytes.NewReader(result), int64(len(result)))
	if err != nil {
		return errors.Wrap(err, "can't read ["+zipFilename+"] as a zip file")
	}

	for _, file := range r.File {
		fPath := path.Join(cfg.RootPath, file.Name)
		if file.FileInfo().IsDir() {
			// Make Folder
			_ = os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		contentBytes, err := ioutil.ReadAll(rc)
		if err != nil {
			return err
		}

		tmpl, err := template.New("bootstrap").Parse(string(contentBytes))
		if err != nil {
			return err
		}
		out := &bytes.Buffer{}
		err = tmpl.Execute(out, cfg)
		if err != nil {
			return err
		}

		_, err = outFile.Write(out.Bytes())
		_ = outFile.Close()
		_ = rc.Close()
		if err != nil {
			return err
		}
	}

	logger.Info(fmt.Sprintf("extracted [%v] to [%v]", prototype.Key, cfg.RootPath))
	return nil
}
