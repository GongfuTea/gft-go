package zs

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftGsZsLqk struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	GsZsLqkData     `bson:",inline" json:",inline"`
}

type GsZsLqkData struct {
	Ksh   string         `bson:"ksh" json:"ksh,omitempty"`
	Xh    string         `bson:"xh" json:"xh,omitempty"`
	Xm    string         `bson:"xm" json:"xm,omitempty"`
	Csrq  *time.Time     `bson:"csrq" json:"csrq,omitempty"`
	Nj    int            `bson:"nj" json:"nj,omitempty"`
	Xz    float64        `bson:"xz" json:"xz,omitempty"`
	Sfzh  string         `bson:"sfzh" json:"sfzh,omitempty"`
	Zjlxm string         `bson:"zjlxm" json:"zjlxm,omitempty"` // 证件类型
	Xbm   string         `bson:"xbm" json:"xbm,omitempty"`     // 性别
	Mzm   string         `bson:"mzm" json:"mzm,omitempty"`     // 民族
	Zzmmm string         `bson:"zzmmm" json:"zzmmm,omitempty"` // 政治面貌
	Yxsm  string         `bson:"yxsm" json:"yxsm,omitempty"`   // 院系所
	Zydm  string         `bson:"zydm" json:"zydm,omitempty"`   // 专业代码
	Zymc  string         `bson:"zymc" json:"zymc,omitempty"`   // 专业名称
	Pyccm string         `bson:"pyccm" json:"pyccm,omitempty"` // 层次：硕士、博士
	Xxxsm string         `bson:"xxxsm" json:"xxxsm,omitempty"` // 学习形式 XX.XXXS
	Pyfsm string         `bson:"pyfsm" json:"pyfsm,omitempty"` // 培养方式：非定向，定向
	Note  string         `bson:"note" json:"note,omitempty"`
	Meta  map[string]any `bson:"meta" json:"meta,omitempty"`
}

func NewGsZsLqk(data GsZsLqkData) *GftGsZsLqk {
	item := &GftGsZsLqk{
		GsZsLqkData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
