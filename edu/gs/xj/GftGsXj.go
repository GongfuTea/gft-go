package xj

import (
	"reflect"
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/GongfuTea/gft-go/x"
	"github.com/GongfuTea/gft-go/x/jsonx"
	"github.com/google/uuid"
)

type GftGsXj struct {
	types.Entity      `bson:",inline" json:",inline"`
	types.ModelBase   `bson:",inline" json:",inline"`
	types.GftTimeline `bson:",inline"`
	GsXjData          `bson:",inline" json:",inline"`
}

type GsXjData struct {
	Ksh     string         `bson:"ksh" json:"ksh,omitempty"`
	Xh      string         `bson:"xh" json:"xh,omitempty"`
	Xm      string         `bson:"xm" json:"xm,omitempty"`
	Csrq    *time.Time     `bson:"csrq" json:"csrq,omitempty"`
	Rxrq    *time.Time     `bson:"rxrq" json:"rxrq,omitempty"`
	Byrq    *time.Time     `bson:"byrq" json:"byrq,omitempty"`
	Yjbyrq  *time.Time     `bson:"yjbyrq" json:"yjbyrq,omitempty"`
	Nj      int            `bson:"nj" json:"nj,omitempty"`
	Xz      float64        `bson:"xz" json:"xz,omitempty"`
	Sfzh    string         `bson:"sfzh" json:"sfzh,omitempty"`
	Zjlx    string         `bson:"zjlx" json:"zjlx,omitempty"`       // 证件类型
	Zjlxm   string         `bson:"zjlxm" json:"zjlxm,omitempty"`     // 证件类型
	Xb      string         `bson:"xb" json:"xb,omitempty"`           // 性别
	Xbm     string         `bson:"xbm" json:"xbm,omitempty"`         // 性别
	Mz      string         `bson:"mz" json:"mz,omitempty"`           // 民族
	Mzm     string         `bson:"mzm" json:"mzm,omitempty"`         // 民族
	Zzmm    string         `bson:"zzmm" json:"zzmm,omitempty"`       // 政治面貌
	Zzmmm   string         `bson:"zzmmm" json:"zzmmm,omitempty"`     // 政治面貌
	Yxs     string         `bson:"yxs" json:"yxs,omitempty"`         // 院系所
	Yxsm    string         `bson:"yxsm" json:"yxsm,omitempty"`       // 院系所
	Zydm    string         `bson:"zydm" json:"zydm,omitempty"`       // 专业代码
	Zymc    string         `bson:"zymc" json:"zymc,omitempty"`       // 专业名称
	Pycc    string         `bson:"pycc" json:"pycc,omitempty"`       // 层次：硕士、博士
	Pyccm   string         `bson:"pyccm" json:"pyccm,omitempty"`     // 层次：硕士、博士
	Xxxs    string         `bson:"xxxs" json:"xxxs,omitempty"`       // 学习形式 XX.XXXS
	Xxxsm   string         `bson:"xxxsm" json:"xxxsm,omitempty"`     // 学习形式 XX.XXXS
	Pyfs    string         `bson:"pyfs" json:"pyfs,omitempty"`       // 培养方式：非定向，定向
	Pyfsm   string         `bson:"pyfsm" json:"pyfsm,omitempty"`     // 培养方式：非定向，定向
	Xsdqzt  string         `bson:"xsdqzt" json:"xsdqzt,omitempty"`   // 学生当前状态
	Xsdqztm string         `bson:"xsdqztm" json:"xsdqztm,omitempty"` // 学生当前状态
	Note    string         `bson:"note" json:"note,omitempty"`
	Meta    map[string]any `bson:"meta" json:"meta,omitempty"`
}

func NewCmsBanner(data GsXjData) *GftGsXj {
	item := &GftGsXj{
		GsXjData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}

func (xj *GftGsXj) UpdateByDiff(diff map[string]types.GftTimelineDiff) {

	el := reflect.ValueOf(xj).Elem()

	for k, v := range diff {
		jsonx.PrintAsJson("diff:"+k, v)

		f := el.FieldByName(k)
		if f.IsValid() {
			if f.CanSet() {
				f.Set(reflect.ValueOf(v[1]))
			}
		}
	}
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
			// println(xj.Xm, f, val1.Interface(), val1.String(), val2.Interface(), val2.String())
			diff[f] = types.GftTimelineDiff{val2.Interface(), val1.Interface()}
			jsonx.PrintAsJson("diff: "+xj.Xm, diff)

			// if val2.String() == "441402197908292352" {
			// jsonx.PrintAsJson("diff", diff)
			// 	jsonx.PrintAsJson("other", other)
			// }
		}
	}

	return diff
}
