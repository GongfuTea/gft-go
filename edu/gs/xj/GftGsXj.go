package xj

import (
	"reflect"
	"time"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
	"github.com/GongfuTea/gft-go/x"
)

type GftGsXj struct {
	db.DbEntity       `bson:",inline"`
	types.GftTimeline `bson:",inline"`
	Ksh               string         `bson:"ksh" json:"ksh,omitempty"`
	Xh                string         `bson:"xh" json:"xh,omitempty"`
	Xm                string         `bson:"xm" json:"xm,omitempty"`
	Csrq              *time.Time     `bson:"csrq" json:"csrq,omitempty"`
	Rxrq              *time.Time     `bson:"rxrq" json:"rxrq,omitempty"`
	Byrq              *time.Time     `bson:"byrq" json:"byrq,omitempty"`
	Yjbyrq            *time.Time     `bson:"yjbyrq" json:"yjbyrq,omitempty"`
	Nj                int            `bson:"nj" json:"nj,omitempty"`
	Xz                float64        `bson:"xz" json:"xz,omitempty"`
	Sfzh              string         `bson:"sfzh" json:"sfzh,omitempty"`
	Zjlxm             string         `bson:"zjlxm" json:"zjlxm,omitempty"`     // 证件类型
	Xb                string         `bson:"xb" json:"xb,omitempty"`           // 性别
	Xbm               string         `bson:"xbm" json:"xbm,omitempty"`         // 性别
	Mz                string         `bson:"mz" json:"mz,omitempty"`           // 民族
	Mzm               string         `bson:"mzm" json:"mzm,omitempty"`         // 民族
	Zzmm              string         `bson:"zzmm" json:"zzmm,omitempty"`       // 政治面貌
	Zzmmm             string         `bson:"zzmmm" json:"zzmmm,omitempty"`     // 政治面貌
	Yxs               string         `bson:"yxs" json:"yxs,omitempty"`         // 院系所
	Yxsm              string         `bson:"yxsm" json:"yxsm,omitempty"`       // 院系所
	Zydm              string         `bson:"zydm" json:"zydm,omitempty"`       // 专业代码
	Zymc              string         `bson:"zymc" json:"zymc,omitempty"`       // 专业名称
	Pycc              string         `bson:"pycc" json:"pycc,omitempty"`       // 层次：硕士、博士
	Pyccm             string         `bson:"pyccm" json:"pyccm,omitempty"`     // 层次：硕士、博士
	Xxxs              string         `bson:"xxxs" json:"xxxs,omitempty"`       // 学习形式 XX.XXXS
	Xxxsm             string         `bson:"xxxsm" json:"xxxsm,omitempty"`     // 学习形式 XX.XXXS
	Pyfs              string         `bson:"pyfs" json:"pyfs,omitempty"`       // 培养方式：非定向，定向
	Pyfsm             string         `bson:"pyfsm" json:"pyfsm,omitempty"`     // 培养方式：非定向，定向
	Xsdqzt            string         `bson:"xsdqzt" json:"xsdqzt,omitempty"`   // 学生当前状态
	Xsdqztm           string         `bson:"xsdqztm" json:"xsdqztm,omitempty"` // 学生当前状态
	Note              string         `bson:"note" json:"note,omitempty"`
	Meta              map[string]any `bson:"meta" json:"meta,omitempty"`
}

func (xj *GftGsXj) Diff(other *GftGsXj) map[string]types.GftTimelineDiff {
	diff := make(map[string]types.GftTimelineDiff, 0)

	ignoreDiffFields := []string{"DbEntity", "GftTimeline"}

	xj1 := reflect.ValueOf(xj)
	xj2 := reflect.ValueOf(other)
	fileds := make([]string, 0)

	t := reflect.TypeOf(xj).Elem()
	c := t.NumField()
	for i := 0; i < c; i++ {
		name := t.Field(i).Name
		if !x.Contains(ignoreDiffFields, name) {
			fileds = append(fileds, name)
		}
	}

	for _, f := range fileds {
		val1 := reflect.Indirect(xj1).FieldByName(f)
		val2 := reflect.Indirect(xj2).FieldByName(f)

		// println("val1,val2: ", val1, val2)
		if !reflect.DeepEqual(val1.Interface(), val2.Interface()) {
			// println("field", f, val1.Interface(), val1.String(), val2.Interface(), val2.String())
			diff[f] = types.GftTimelineDiff{val1.Interface(), val2.Interface()}

			// if val2.String() == "441402197908292352" {
			// 	jsonx.PrintAsJson("xj", xj)
			// 	jsonx.PrintAsJson("other", other)
			// }
		}
	}

	return diff
}

type GftGsXjList []*GftGsXj

func (l GftGsXjList) Includes(it *GftGsXj) bool {
	for _, x := range l {
		if x.Xh == it.Xh {
			return true
		}
	}

	return false
}

func (l GftGsXjList) FindByXh(xh string) *GftGsXj {
	for _, x := range l {
		if x.Xh == xh {
			return x
		}
	}
	return nil
}

func (l GftGsXjList) FindAdded(oldList GftGsXjList) (added GftGsXjList) {
	added = make(GftGsXjList, 0)

	for _, x := range l {
		if !oldList.Includes(x) {
			added = append(added, x)
		}
	}

	return added
}

func (l GftGsXjList) FindRemoved(oldList GftGsXjList) GftGsXjList {
	return oldList.FindAdded(l)
}

func (l GftGsXjList) FindUpdated(oldList GftGsXjList) (updated GftGsXjList) {
	updated = make(GftGsXjList, 0)

	for _, x := range l {
		if old := oldList.FindByXh(x.Xh); old != nil {
			diff := x.Diff(old)
			if len(diff) > 0 {
				x.TlDiff = diff
				updated = append(updated, x)
			}
		}
	}

	return updated
}
