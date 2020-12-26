package xls

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Set the column headers for the provided excelize.File
func SetColumnHeaders(key string, cols []string, f *excelize.File) {
	c := 'A'
	for _, col := range cols {
		f.SetCellValue(key, string(c)+"1", col)
		c++
	}
}

// Set the full data for the provided excelize.File
func SetData(key string, firstRow int, data [][]interface{}, f *excelize.File) {
	for rowIdx, row := range data {
		for colIdx, col := range row {
			f.SetCellValue(key, fmt.Sprintf("%v%v", string(rune(colIdx+'A')), rowIdx+firstRow), col)
		}
	}
}

// Set the title of the first sheet in the provided excelize.File
func SetFirstSheetTitle(t string, f *excelize.File) {
	f.SetSheetName(DefSheet, t)
}

// Sets column widths for the first sheet in the provided excelize.File
func SetColumnWidths(key string, widths []int, f *excelize.File) {
	for i, w := range widths {
		col := string(rune('A' + i))
		f.SetColWidth(key, col, col, float64(w))
	}
}
