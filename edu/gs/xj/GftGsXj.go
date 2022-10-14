package xj

import (
	"time"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
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
