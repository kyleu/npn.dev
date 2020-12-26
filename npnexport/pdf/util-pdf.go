package pdf

import (
	"github.com/johnfercher/maroto/pkg/consts"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// Creates a table row in the provided PDF file
func TR(content func(), height float64, m pdfgen.Maroto) {
	m.Row(height, content)
}

// Creates a column in the provided PDF file
func Col(content func(), width uint, m pdfgen.Maroto) {
	m.Col(width, content)
}

// Creates a table header in the provided PDF file
func TH(content string, width uint, m pdfgen.Maroto) {
	Col(func() { m.Text(content, props.Text{Style: consts.Bold}) }, width, m)
}

// Creates a table cell in the provided PDF file
func TD(content string, width uint, m pdfgen.Maroto) {
	Col(func() { m.Text(content) }, width, m)
}

// Creates a table in the provided PDF file
func Table(cols []string, data [][]string, sizes []uint, m pdfgen.Maroto) {
	m.TableList(cols, data, props.TableList{
		HeaderProp:         props.TableListContent{Style: consts.Bold, GridSizes: sizes},
		ContentProp:        props.TableListContent{GridSizes: sizes},
		HeaderContentSpace: 2,
	})
}

// Creates a horizontal rule in the provided PDF file
func HR(m pdfgen.Maroto) {
	m.Line(12)
}

// Creates a caption in the provided PDF file
func Caption(title string, m pdfgen.Maroto) {
	TR(func() { Col(func() { m.Text(title, props.Text{Size: 12}) }, 12, m) }, 10, m)
}

// Creates a full table row in the provided PDF file
func DetailRow(k string, v string, m pdfgen.Maroto) {
	TR(func() {
		TH(k, 2, m)
		TD(v, 10, m)
	}, 8, m)
}
