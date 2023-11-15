package xj

import (
	"time"

	"github.com/GongfuTea/gft-go/base"
	bmgo "github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/mgo"
	"github.com/GongfuTea/gft-go/x/timex"
	"github.com/GongfuTea/gft-go/x/xlsx"
)

type XxwFile struct {
	*xlsx.Xlsx
	Items []xlsx.ItemMap

	XjList     GftGsXjList
	YxsList    []*jc.GftGsYxs
	MzList     []*base.GftDictItem
	ZzmmList   []*base.GftDictItem
	ZjlxList   []*base.GftDictItem
	XxxsList   []*base.GftDictItem
	PyfsList   []*base.GftDictItem
	XsdqztList []*base.GftDictItem
}

func NewXxwFile() *XxwFile {
	xxw := &XxwFile{}
	xxw.YxsList, _ = mgo.GsYxsRepo.All()
	xxw.MzList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.MZ")
	xxw.ZzmmList, _ = bmgo.DictItemRepo.FindByCategoryId("GB.ZZMM")
	xxw.ZjlxList, _ = bmgo.DictItemRepo.FindByCategoryId("JY.SFZJLX")
	xxw.XxxsList, _ = bmgo.DictItemRepo.FindByCategoryId("XX.XXXS")
	xxw.PyfsList, _ = bmgo.DictItemRepo.FindByCategoryId("JY.PYFS")
	xxw.XsdqztList, _ = bmgo.DictItemRepo.FindByCategoryId("JY.XSDQZT")
	return xxw
}

func (file *XxwFile) LoadFile(filename string) *XxwFile {
	file.Xlsx = xlsx.OpenFile(filename)
	file.Xlsx.SetSheet(file.Xlsx.GetSheetName(0))
	file.Items = file.GetRecords(&xlsx.GetRecordOptions{LowerHeaderName: true})
	file.XjList = GftGsXjList{}

	for _, row := range file.Items {
		stu := file.getXj(row)
		file.XjList = append(file.XjList, stu)
	}

	return file
}

func (x *XxwFile) Xbm(it xlsx.ItemMap) string {
	xb := it.GetField("xb")
	if xb == "男" {
		return "1"
	} else if xb == "女" {
		return "2"
	} else {
		return "0"
	}
}

func (x *XxwFile) GetTime(it xlsx.ItemMap, field string) *time.Time {
	if str := it.GetField(field); str == "" {
		return nil
	} else {
		v := timex.ParseLocalUTC(str)
		return &v
	}
}

func (x *XxwFile) GetZzmm(it xlsx.ItemMap) *base.GftDictItem {
	val := it.GetFieldDefault("zzmm", "群众")

	for _, it := range x.ZzmmList {
		if it.Name == val || it.Nickname == val {
			return it
		}
	}
	panic("Error:" + val)
}

func (x *XxwFile) GetMz(it xlsx.ItemMap) *base.GftDictItem {
	val := it.GetField("mz")
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

func (x *XxwFile) GetZjlx(it xlsx.ItemMap) *base.GftDictItem {
	val := it.GetField("sfzh")
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

func (x *XxwFile) GetYxs(it xlsx.ItemMap) *jc.GftGsYxs {
	val := it.GetField("fy")
	if val == "长江新闻与传播学" {
		val = "长江新闻与传播学院"
	}
	if val == "社科部" {
		val = "马克思主义学院"
	}
	for _, it := range x.YxsList {
		if it.Name == val || it.Nickname == val {
			return it
		}
	}
	panic("Error:" + val)
}

func (x *XxwFile) GetPyccm(it xlsx.ItemMap) string {
	pycc := it.GetField("cc")
	if pycc == "硕士研究生" {
		return "2"
	} else if pycc == "博士研究生" {
		return "1"
	} else {
		return "9"
	}
}

func (x *XxwFile) GetXxxs(it xlsx.ItemMap) *base.GftDictItem {
	val := it.GetField("xxxs")
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

func (x *XxwFile) GetXjzt(it xlsx.ItemMap) *base.GftDictItem {
	val := it.GetField("zczt")
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

func (x *XxwFile) GetPyfs(it xlsx.ItemMap) *base.GftDictItem {
	val := it.GetField("sx")
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

func (file *XxwFile) getXj(item xlsx.ItemMap) *GftGsXj {
	rxrq := file.GetTime(item, "rxrq")

	stu := &GftGsXj{
		Ksh:     item.GetField("ksh"),
		Xh:      item.GetField("xh"),
		Xm:      item.GetField("xm"),
		Xb:      item.GetField("xb"),
		Xbm:     file.Xbm(item),
		Csrq:    file.GetTime(item, "csrq"),
		Sfzh:    item.GetField("xm"),
		Zzmmm:   file.GetZzmm(item).Code,
		Zzmm:    file.GetZzmm(item).Name,
		Mz:      file.GetMz(item).Name,
		Mzm:     file.GetMz(item).Code,
		Zjlx:    file.GetZjlx(item).Name,
		Zjlxm:   file.GetZjlx(item).Code,
		Zydm:    item.GetField("zydm"),
		Zymc:    item.GetField("zymc"),
		Yxsm:    file.GetYxs(item).Code,
		Yxs:     file.GetYxs(item).Name,
		Pycc:    item.GetField("cc"),
		Pyccm:   file.GetPyccm(item),
		Xz:      item.GetFloat("xz"),
		Xxxs:    file.GetXxxs(item).Name,
		Xxxsm:   file.GetXxxs(item).Code,
		Nj:      item.GetInt("dqszj"),
		Rxrq:    rxrq,
		Yjbyrq:  file.GetTime(item, "yjbyrq"),
		Xsdqztm: file.GetXjzt(item).Code,
		Xsdqzt:  file.GetXjzt(item).Name,
		Pyfs:    file.GetPyfs(item).Name,
		Pyfsm:   file.GetPyfs(item).Code,
	}

	return stu
}
