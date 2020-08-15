package pdf

import (
	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/johnfercher/maroto/pkg/consts"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func Render(rsp interface{}, url string, callback func(rsp interface{}, m pdfgen.Maroto) (string, error)) (string, []byte, error) {
	m := newDoc()
	writeDocHeader(url, m)
	filename, err := callback(rsp, m)
	if err != nil {
		return filename, nil, err
	}
	return response(filename, m)
}

func newDoc() pdfgen.Maroto {
	m := pdfgen.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)
	return m
}

func writeDocHeader(url string, m pdfgen.Maroto) {
	m.RegisterHeader(func() {
		TR(func() {
			Col(func() {
				m.Text(npncore.AppName, props.Text{Size: 16, Align: consts.Left})
				m.Text(url, props.Text{Size: 8, Align: consts.Right})
			}, 12, m)
		}, 10, m)
	})
}

func response(fn string, m pdfgen.Maroto) (string, []byte, error) {
	buff, err := m.Output()
	if err != nil {
		return fn, nil, errors.Wrap(err, "error writing PDF output")
	}

	if len(fn) == 0 {
		fn = npncore.KeyExport
	}
	return fn, buff.Bytes(), nil
}
