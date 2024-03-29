package xlsx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	*excelize.File

	Filename string
	CurSheet string
}

type GetRecordOptions struct {
	LowerHeaderName bool
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

func (x *Xlsx) SetCellVal(row int, col string, val any) error {
	axis := fmt.Sprintf("%s%d", col, row)
	return x.SetCellValAxis(axis, val)
}

func (x *Xlsx) SetCellValAxis(axis string, val any) error {
	return x.SetCellValue(x.CurSheet, axis, val)
}

func (x *Xlsx) GetCellStr(col string, row int) (string, error) {
	return x.GetCellValue(x.CurSheet, fmt.Sprintf("%s%d", col, row))
}

func (x *Xlsx) GetCellStr2(col string, row int) string {
	txt, _ := x.GetCellStr(col, row)
	return strings.TrimSpace((txt))
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

func (x *Xlsx) TotalRows() int {
	if rows, err := x.GetRows(x.CurSheet); err == nil {
		return len(rows)
	}
	return 0
}

func (x *Xlsx) GetRecords(opt *GetRecordOptions) (items []ItemMap) {
	items = make([]ItemMap, 0)
	rows, _ := x.GetRows(x.CurSheet, excelize.Options{RawCellValue: true})
	cols := rows[0]
	if opt != nil && opt.LowerHeaderName {
		for i, c := range cols {
			cols[i] = strings.ToLower(c)
		}
	}

	for _, row := range rows[1:] {
		item := make(ItemMap, len(cols))
		rowLen := len(row)
		for c, col := range cols {
			val := ""
			if rowLen > c {
				val = row[c]
			}
			item[col] = val
		}
		items = append(items, item)
	}

	return items
}
