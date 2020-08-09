package bootstrap

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path"

	"emperror.dev/errors"
)

func PersistAll() error {
	for _, project := range AllPrototypes {
		err := Persist(project)
		if err != nil {
			return errors.Wrap(err, "error persisting project ["+project.Key+"]")
		}
	}
	return nil
}

func Persist(prototype *Prototype) error {
	outFilename := "web/assets/bootstrap/" + prototype.Key + ".zip"
	outFile, err := os.Create(outFilename)
	if err != nil {
		return errors.Wrap(err, "cannot create bootstrap project file ["+outFilename+"]")
	}
	defer func() { _ = outFile.Close() }()

	zw := zip.NewWriter(outFile)
	for _, folder := range prototype.Folders {
		err = addFolder(zw, path.Join("bootstrap", folder), "")
	}
	if err != nil {
		return errors.Wrap(err, "cannot add files to bootstrap project ["+prototype.Key+"]")
	}
	return zw.Close()
}

func addFolder(zw *zip.Writer, folder string, root string) error {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = processFile(zw, root, folder, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func processFile(zw *zip.Writer, root string, folder string, file os.FileInfo) error {
	newFilename := path.Join(folder, file.Name())
	if file.IsDir() {
		newRoot := path.Join(root, file.Name())
		err := addFolder(zw, newFilename, newRoot)
		return err
	}
	dat, err := ioutil.ReadFile(newFilename)
	if err != nil {
		return errors.Wrap(err, "can't read file ["+file.Name()+"]")
	}

	// Add some files to the archive.
	f, err := zw.Create(path.Join(root, file.Name()))
	if err != nil {
		return errors.Wrap(err, "can't create zip entry for ["+path.Join(root, file.Name())+"]")
	}
	_, err = f.Write(dat)
	return errors.Wrap(err, "can't write zip entry for ["+path.Join(root, file.Name())+"]")
}
