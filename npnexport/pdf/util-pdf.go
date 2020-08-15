package pdf

import (
	"github.com/johnfercher/maroto/pkg/consts"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func TR(content func(), height float64, m pdfgen.Maroto) {
	m.Row(height, content)
}

func Col(content func(), width uint, m pdfgen.Maroto) {
	m.Col(width, content)
}

func TH(content string, width uint, m pdfgen.Maroto) {
	Col(func() { m.Text(content, props.Text{Style: consts.Bold}) }, width, m)
}

func TD(content string, width uint, m pdfgen.Maroto) {
	Col(func() { m.Text(content) }, width, m)
}

func Table(cols []string, data [][]string, sizes []uint, m pdfgen.Maroto) {
	m.TableList(cols, data, props.TableList{
		HeaderProp:         props.TableListContent{Style: consts.Bold, GridSizes: sizes},
		ContentProp:        props.TableListContent{GridSizes: sizes},
		HeaderContentSpace: 2,
	})
}

func HR(m pdfgen.Maroto) {
	m.Line(12)
}

func Caption(title string, m pdfgen.Maroto) {
	TR(func() { Col(func() { m.Text(title, props.Text{Size: 12}) }, 12, m) }, 10, m)
}

func DetailRow(k string, v string, m pdfgen.Maroto) {
	TR(func() {
		TH(k, 2, m)
		TD(v, 10, m)
	}, 8, m)
}
