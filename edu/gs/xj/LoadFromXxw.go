package xj

import (
	"strconv"
	"time"

	"github.com/GongfuTea/gft-go/base"
	bmgo "github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/mgo"
	"github.com/GongfuTea/gft-go/x/timex"
	"github.com/GongfuTea/gft-go/x/xlsx"
	"github.com/xuri/excelize/v2"
)

// * 从学信网名单加载学籍名单
type LoadFromXxw struct {
	*xlsx.Xlsx

	XjList     GftGsXjList
	YxsList    []*jc.GftGsYxs
	MzList     []*base.GftDictItem
	ZzmmList   []*base.GftDictItem
	ZjlxList   []*base.GftDictItem
	XxxsList   []*base.GftDictItem
	PyfsList   []*base.GftDictItem
	XsdqztList []*base.GftDictItem
}

func (file *LoadFromXxw) Load(filename string) *LoadFromXxw {
	file.Xlsx = xlsx.OpenFile(filename)
	file.Xlsx.SetSheet(file.Xlsx.GetSheetName(0))
	file.YxsList, _ = mgo.GsYxsRepo.All()
	file.MzList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.MZ")
	file.ZzmmList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.ZZMM")
	file.ZjlxList, _ = bmgo.DictItemRepo.FindByCategoryId("JY.SFZJLX")
	file.XxxsList, _ = bmgo.DictItemRepo.FindByCategoryId("XX.XXXS")
	file.PyfsList, _ = bmgo.DictItemRepo.FindByCategoryId("JY.PYFS")
	file.XsdqztList, _ = bmgo.DictItemRepo.FindByCategoryId("JY.XSDQZT")
	file.XjList = GftGsXjList{}

	rows, _ := file.Xlsx.GetRows(file.Xlsx.CurSheet, excelize.Options{RawCellValue: true})

	for _, row := range rows[1:] {
		rxrq := file.GetTime(row[20])
		stu := &GftGsXj{
			Ksh:     row[0],
			Xh:      row[2],
			Xm:      row[3],
			Xb:      row[4],
			Xbm:     file.Xbm(row[4]),
			Csrq:    file.GetTime(row[5]),
			Sfzh:    row[6],
			Zzmmm:   file.GetZzmm(row[7]).Code,
			Zzmm:    file.GetZzmm(row[7]).Name,
			Mz:      file.GetMz(row[8]).Name,
			Mzm:     file.GetMz(row[8]).Code,
			Zjlx:    file.GetZjlx(row[6]).Name,
			Zjlxm:   file.GetZjlx(row[6]).Code,
			Zydm:    row[11],
			Zymc:    row[12],
			Yxsm:    file.GetYxs(row[13]).Code,
			Yxs:     file.GetYxs(row[13]).Name,
			Pycc:    row[16],
			Pyccm:   file.GetPyccm(row[16]),
			Xz:      file.GetFloat(row[17]),
			Xxxs:    file.GetXxxs(row[18]).Name,
			Xxxsm:   file.GetXxxs(row[18]).Code,
			Nj:      file.GetInt(row[19]),
			Rxrq:    rxrq,
			Yjbyrq:  file.GetTime(row[21]),
			Xsdqztm: file.GetXjzt(row[22]).Code,
			Xsdqzt:  file.GetXjzt(row[22]).Name,
			Pyfs:    file.GetPyfs(row[23]).Name,
			Pyfsm:   file.GetPyfs(row[23]).Code,
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

func (x *LoadFromXxw) GetPyfs(val string) *base.GftDictItem {
	code := "11"
	if val == "定向" {
		code = "12"
	} else if val == "委培" {
		code = "12"
	}
	for _, it := range x.PyfsList {
		if it.Code == code {
			return it
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetXxxs(val string) *base.GftDictItem {
	name := val
	if val != "全日制" && val != "非全日制" {
		name = "全日制"
	}

	for _, it := range x.XxxsList {
		if it.Name == name {
			return it
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetXjzt(val string) *base.GftDictItem {
	code := "01"
	if val == "注册学籍" {
		code = "01"
	} else if val == "休学" {
		code = "02"
	} else if val == "暂缓注册" {
		code = "12"
	} else if val == "待注册" {
		code = "12"
	} else {
		panic("Error:" + val)
	}

	for _, it := range x.XsdqztList {
		if it.Code == code {
			return it
		}
	}
	panic("Error:" + val)

}
func (x *LoadFromXxw) GetYxs(val string) *jc.GftGsYxs {
	for _, it := range x.YxsList {
		if it.Name == val || it.Nickname == val {
			return it
		}
	}
	panic("Error:" + val)
}
func (x *LoadFromXxw) GetYxsm(val string) string {
	return x.GetYxs(val).Code
}

func (x *LoadFromXxw) GetMz(val string) *base.GftDictItem {
	if val == "其他" {
		val = "其它"
	}
	for _, it := range x.MzList {
		if it.Name == val {
			return it
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetMzm(val string) string {
	return x.GetMz(val).Code
}

func (x *LoadFromXxw) GetZzmm(val string) *base.GftDictItem {
	for _, it := range x.ZzmmList {
		if it.Name == val || it.Nickname == val {
			return it
		}
	}
	panic("Error:" + val)
}

func (x *LoadFromXxw) GetZzmmm(val string) string {
	return x.GetZzmm(val).Code
}

func (x *LoadFromXxw) GetZjlx(val string) *base.GftDictItem {
	code := "1"
	if len(val) == 18 {
		code = "1"
	}
	if val == "M427772(A)" || val == "M614972(9)" || val == "M899175(3)" {
		code = "6"
	}
	if val == "F229619072" || val == "L125494231" {
		code = "8"
	}
	for _, it := range x.ZjlxList {
		if it.Code == code {
			return it
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
	v := timex.ParseLocalUTC(str)
	return &v
}
