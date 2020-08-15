package xls

import (
	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/360EntSecGroup-Skylar/excelize"
)

var DefSheet = "Sheet1"

func Render(rsp interface{}, url string, callback func(rsp interface{}, f *excelize.File) (string, string, error)) (string, []byte, error) {
	f := newDoc()
	filename, title, err := callback(rsp, f)
	if err != nil {
		return filename, nil, err
	}
	writeAboutSheet(url, title, f)
	SetFirstSheetTitle(filename, f)
	return response(filename, f)
}

func newDoc() *excelize.File {
	f := excelize.NewFile()
	f.SetActiveSheet(1)
	return f
}

func writeAboutSheet(url string, title string, f *excelize.File) {
	key := npncore.AppName
	f.NewSheet(key)
	f.SetCellValue(key, "A1", npncore.AppName)
	f.SetCellValue(key, "A2", url)
	f.SetCellValue(key, "A3", title)
	SetColumnWidths(key, []int{64}, f)
}

func response(fn string, f *excelize.File) (string, []byte, error) {
	buff, err := f.WriteToBuffer()
	if err != nil {
		return fn, nil, errors.Wrap(err, "error writing PDF output")
	}

	if len(fn) == 0 {
		fn = npncore.KeyExport
	}
	return fn, buff.Bytes(), nil
}
