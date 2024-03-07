package commands

import "time"

type SaveGsXj struct {
	Id    string        `json:"id,omitempty"`
	Input SaveGsXjInput `json:"input"`
}

type DelGsXj struct {
	Id string `json:"id"`
}

type SaveGsXjInput struct {
	Ksh     string         `bson:"ksh,omitempty" json:"ksh,omitempty"`
	Xh      string         `bson:"xh,omitempty" json:"xh,omitempty"`
	Xm      string         `bson:"xm,omitempty" json:"xm,omitempty"`
	Csrq    *time.Time     `bson:"csrq,omitempty" json:"csrq,omitempty"`
	Rxrq    *time.Time     `bson:"rxrq,omitempty" json:"rxrq,omitempty"`
	Byrq    *time.Time     `bson:"byrq,omitempty" json:"byrq,omitempty"`
	Yjbyrq  *time.Time     `bson:"yjbyrq,omitempty" json:"yjbyrq,omitempty"`
	Nj      int            `bson:"nj,omitempty" json:"nj,omitempty"`
	Xz      float64        `bson:"xz,omitempty" json:"xz,omitempty"`
	Sfzh    string         `bson:"sfzh,omitempty" json:"sfzh,omitempty"`
	Zjlx    string         `bson:"zjlx,omitempty" json:"zjlx,omitempty"`       // 证件类型
	Zjlxm   string         `bson:"zjlxm,omitempty" json:"zjlxm,omitempty"`     // 证件类型
	Xb      string         `bson:"xb,omitempty" json:"xb,omitempty"`           // 性别
	Xbm     string         `bson:"xbm,omitempty" json:"xbm,omitempty"`         // 性别
	Mz      string         `bson:"mz,omitempty" json:"mz,omitempty"`           // 民族
	Mzm     string         `bson:"mzm,omitempty" json:"mzm,omitempty"`         // 民族
	Zzmm    string         `bson:"zzmm,omitempty" json:"zzmm,omitempty"`       // 政治面貌
	Zzmmm   string         `bson:"zzmmm,omitempty" json:"zzmmm,omitempty"`     // 政治面貌
	Yxs     string         `bson:"yxs,omitempty" json:"yxs,omitempty"`         // 院系所
	Yxsm    string         `bson:"yxsm,omitempty" json:"yxsm,omitempty"`       // 院系所
	Zydm    string         `bson:"zydm,omitempty" json:"zydm,omitempty"`       // 专业代码
	Zymc    string         `bson:"zymc,omitempty" json:"zymc,omitempty"`       // 专业名称
	Pycc    string         `bson:"pycc,omitempty" json:"pycc,omitempty"`       // 层次：硕士、博士
	Pyccm   string         `bson:"pyccm,omitempty" json:"pyccm,omitempty"`     // 层次：硕士、博士
	Xxxs    string         `bson:"xxxs,omitempty" json:"xxxs,omitempty"`       // 学习形式 XX.XXXS
	Xxxsm   string         `bson:"xxxsm,omitempty" json:"xxxsm,omitempty"`     // 学习形式 XX.XXXS
	Pyfs    string         `bson:"pyfs,omitempty" json:"pyfs,omitempty"`       // 培养方式：非定向，定向
	Pyfsm   string         `bson:"pyfsm,omitempty" json:"pyfsm,omitempty"`     // 培养方式：非定向，定向
	Xsdqzt  string         `bson:"xsdqzt,omitempty" json:"xsdqzt,omitempty"`   // 学生当前状态
	Xsdqztm string         `bson:"xsdqztm,omitempty" json:"xsdqztm,omitempty"` // 学生当前状态
	Note    string         `bson:"note,omitempty" json:"note,omitempty"`
	Meta    map[string]any `bson:"meta,omitempty" json:"meta,omitempty"`
}
