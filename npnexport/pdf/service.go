package pdf

import (
	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// Renders a PDF document, returning the filename and a byte array representing the contents
func Render(rsp interface{}, url string, callback func(rsp interface{}, m pdf.Maroto) (string, error)) (string, []byte, error) {
	m := newDoc()
	writeDocHeader(url, m)
	filename, err := callback(rsp, m)
	if err != nil {
		return filename, nil, err
	}
	return response(filename, m)
}

func newDoc() pdf.Maroto {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)
	return m
}

func writeDocHeader(url string, m pdf.Maroto) {
	m.RegisterHeader(func() {
		TR(func() {
			Col(func() {
				m.Text(npncore.AppName, props.Text{Size: 16, Align: consts.Left})
				m.Text(url, props.Text{Size: 8, Align: consts.Right})
			}, 12, m)
		}, 10, m)
	})
}

func response(fn string, m pdf.Maroto) (string, []byte, error) {
	buff, err := m.Output()
	if err != nil {
		return fn, nil, errors.Wrap(err, "error writing PDF output")
	}

	if len(fn) == 0 {
		fn = npncore.KeyExport
	}
	return fn, buff.Bytes(), nil
}
