package xj

import (
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
)

type GftGsXj struct {
	*db.DbEntity `bson:",inline"`
	Xh           string            `bson:"xh" json:"xh,omitempty"`
	Xm           string            `bson:"xm" json:"xm,omitempty"`
	Nj           int               `bson:"nj" json:"nj,omitempty"`
	Xz           float32           `bson:"xz" json:"xz,omitempty"`
	Sfzh         string            `bson:"sfzh" json:"sfzh,omitempty"`
	Sfz          types.GftCodeName `bson:"sfz" json:"sfz,omitempty"`   // 身份证
	Xb           types.GftCodeName `bson:"xb" json:"xb,omitempty"`     // 性别
	Mz           types.GftCodeName `bson:"mz" json:"mz,omitempty"`     // 民族
	Zzmm         types.GftCodeName `bson:"zzmm" json:"zzmm,omitempty"` // 政治面貌
	Yx           types.GftCodeName `bson:"yx" json:"yx,omitempty"`     // 院校
	Yxs          types.GftCodeName `bson:"yxs" json:"yxs,omitempty"`   // 院系所
	Zy           types.GftCodeName `bson:"zy" json:"zy,omitempty"`     // 专业
	Cc           types.GftCodeName `bson:"cc" json:"cc,omitempty"`     // 层次：硕士、博士

	Note string            `bson:"note" json:"note,omitempty"`
	Tl   types.GftTimeline `bson:"tl" json:"tl,omitempty"`
}
