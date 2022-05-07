package xlsx

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	*excelize.File

	Filename string
	CurSheet string
}

func OpenFile(filename string) *Xlsx {
	x := &Xlsx{
		Filename: filename,
	}
	x.File, _ = excelize.OpenFile(filename)

	return x
}

func (x *Xlsx) SetSheet(sheet string) *Xlsx {
	x.CurSheet = sheet
	return x
}

func (x *Xlsx) GetCellStr(col string, row int) (string, error) {
	return x.GetCellValue(x.CurSheet, fmt.Sprintf("%s%d", col, row))
}

func (x *Xlsx) GetCellStr2(col string, row int) string {
	txt, _ := x.GetCellStr(col, row)
	return txt
}

func (x *Xlsx) GetCellInt(col string, row int) (int, error) {
	val, _ := x.GetCellValue(x.CurSheet, fmt.Sprintf("%s%d", col, row))
	return strconv.Atoi(val)
}

func (x *Xlsx) GetCellInt2(col string, row int) int {
	val, _ := x.GetCellInt(col, row)
	return val
}

func (x *Xlsx) GetCellFloat(col string, row int) (float64, error) {
	val, _ := x.GetCellValue(x.CurSheet, fmt.Sprintf("%s%d", col, row))
	return strconv.ParseFloat(val, 32)
}

func (x *Xlsx) GetCellFloat2(col string, row int) float64 {
	val, _ := x.GetCellFloat(col, row)
	return val
}
