package xj

import (
	"strconv"
	"time"

	"github.com/GongfuTea/gft-go/base"
	bmgo "github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/mgo"
	"github.com/GongfuTea/gft-go/x/xlsx"
	"github.com/xuri/excelize/v2"
)

// * 从学信网名单加载学籍名单
type LoadFromXxw struct {
	*xlsx.Xlsx

	XjList   GftGsXjList
	YxsList  []*jc.GftGsYxs
	MzList   []*base.GftDictItem
	ZzmmList []*base.GftDictItem
	ZjlxList []*base.GftDictItem
}

func (file *LoadFromXxw) Load(filename string) *LoadFromXxw {
	file.Xlsx = xlsx.OpenFile(filename)
	file.Xlsx.SetSheet(file.Xlsx.GetSheetName(0))
	file.YxsList, _ = mgo.GsYxsRepo.All()
	file.MzList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.MZ")
	file.ZzmmList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.ZZMM")
	file.ZjlxList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.ZJLX")
	file.XjList = GftGsXjList{}

	rows, _ := file.Xlsx.GetRows(file.Xlsx.CurSheet, excelize.Options{RawCellValue: true})

	for _, row := range rows[1:] {
		rxrq := file.GetTime(row[20])
		stu := &GftGsXj{
			Ksh:     row[0],
			Xh:      row[2],
			Xm:      row[3],
			Xbm:     file.Xbm(row[4]),
			Csrq:    file.GetTime(row[5]),
			Sfzh:    row[6],
			Zzmmm:   file.GetZzmmm(row[7]),
			Mzm:     file.GetMzm(row[8]),
			Zjlxm:   file.GetZjlxm(row[6]),
			Zydm:    row[11],
			Zymc:    row[12],
			Yxsm:    file.GetYxsm(row[13]),
			Pyccm:   file.GetPyccm(row[16]),
			Xz:      file.GetFloat(row[17]),
			Xxxsm:   file.GetXxxsm(row[18]),
			Nj:      file.GetInt(row[19]),
			Rxrq:    rxrq,
			Yjbyrq:  file.GetTime(row[21]),
			Xsdqztm: file.GetXjztm(row[22]),
			Pyfsm:   file.GetPyfsm(row[23]),
		}
		file.XjList = append(file.XjList, stu)
	}

	return file
}

func (x *LoadFromXxw) Xbm(xb string) string {
	if xb == "男" {
		return "1"
	} else if xb == "女" {
		return "2"
	} else {
		return "0"
	}
}

func (x *LoadFromXxw) GetPyccm(pycc string) string {
	if pycc == "硕士研究生" {
		return "2"
	} else if pycc == "博士研究生" {
		return "1"
	} else {
		return "9"
	}
}

func (x *LoadFromXxw) GetPyfsm(val string) string {
	if val == "定向" {
		return "12"
	} else if val == "委培" {
		return "22"
	} else {
		return "11"
	}
}

func (x *LoadFromXxw) GetXxxsm(val string) string {
	if val == "全日制" {
		return "1"
	} else if val == "非全日制" {
		return "2"
	} else {
		return "1"
	}
}

func (x *LoadFromXxw) GetXjztm(val string) string {
	if val == "休学" {
		return "02"
	} else if val == "暂缓注册" {
		return "12"
	} else if val == "待注册" {
		return "12"
	} else if val == "注册学籍" {
		return "01"
	}
	panic("Error:" + val)

}

func (x *LoadFromXxw) GetYxsm(val string) string {
	for _, it := range x.YxsList {
		if it.Name == val || it.Nickname == val {
			return it.Code
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetMzm(val string) string {
	if val == "其他" {
		val = "其它"
	}
	for _, it := range x.MzList {
		if it.Name == val {
			return it.Code
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetZzmmm(val string) string {
	for _, it := range x.ZzmmList {
		if it.Name == val || it.Nickname == val {
			return it.Code
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetZjlxm(val string) string {
	if len(val) == 18 {
		return "1"
	}
	if val == "M427772(A)" || val == "M614972(9)" || val == "M899175(3)" {
		return "6"
	}
	if val == "F229619072" || val == "L125494231" {
		return "8"
	}
	for _, it := range x.ZjlxList {
		if it.Name == val {
			return it.Code
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetFloat(str string) float64 {
	v, _ := strconv.ParseFloat(str, 32)
	return v
}

func (x *LoadFromXxw) GetInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func (x *LoadFromXxw) GetTime(str string) *time.Time {
	if str == "" {
		return nil
	}
	v, _ := time.Parse("20060102", str)
	return &v
}
