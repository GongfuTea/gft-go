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
	Ksh     string         `json:"ksh,omitempty"`
	Xh      string         `json:"xh,omitempty"`
	Xm      string         `json:"xm,omitempty"`
	Csrq    *time.Time     `json:"csrq,omitempty"`
	Rxrq    *time.Time     `json:"rxrq,omitempty"`
	Byrq    *time.Time     `json:"byrq,omitempty"`
	Yjbyrq  *time.Time     `json:"yjbyrq,omitempty"`
	Nj      int            `json:"nj,omitempty"`
	Xz      float64        `json:"xz,omitempty"`
	Sfzh    string         `json:"sfzh,omitempty"`
	Zjlx    string         `json:"zjlx,omitempty"`    // 证件类型
	Zjlxm   string         `json:"zjlxm,omitempty"`   // 证件类型
	Xb      string         `json:"xb,omitempty"`      // 性别
	Xbm     string         `json:"xbm,omitempty"`     // 性别
	Mz      string         `json:"mz,omitempty"`      // 民族
	Mzm     string         `json:"mzm,omitempty"`     // 民族
	Zzmm    string         `json:"zzmm,omitempty"`    // 政治面貌
	Zzmmm   string         `json:"zzmmm,omitempty"`   // 政治面貌
	Yxs     string         `json:"yxs,omitempty"`     // 院系所
	Yxsm    string         `json:"yxsm,omitempty"`    // 院系所
	Zydm    string         `json:"zydm,omitempty"`    // 专业代码
	Zymc    string         `json:"zymc,omitempty"`    // 专业名称
	Pycc    string         `json:"pycc,omitempty"`    // 层次：硕士、博士
	Pyccm   string         `json:"pyccm,omitempty"`   // 层次：硕士、博士
	Xxxs    string         `json:"xxxs,omitempty"`    // 学习形式 XX.XXXS
	Xxxsm   string         `json:"xxxsm,omitempty"`   // 学习形式 XX.XXXS
	Pyfs    string         `json:"pyfs,omitempty"`    // 培养方式：非定向，定向
	Pyfsm   string         `json:"pyfsm,omitempty"`   // 培养方式：非定向，定向
	Xsdqzt  string         `json:"xsdqzt,omitempty"`  // 学生当前状态
	Xsdqztm string         `json:"xsdqztm,omitempty"` // 学生当前状态
	Note    string         `json:"note,omitempty"`
	Meta    map[string]any `json:"meta,omitempty"`
}
