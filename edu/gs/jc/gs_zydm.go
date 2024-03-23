package jc

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftGsZydm struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	GsZydmData      `bson:",inline" json:",inline"`
}

type GsZydmData struct {
	Code      string  `bson:"code" json:"code"`
	Name      string  `bson:"name" json:"name"`
	Level     int     `bson:"level" json:"level"`
	Note      string  `bson:"note" json:"note"`
	Xwlxm     string  `bson:"xwlxm" json:"xwlxm"` // 学位类型 xs/zx
	Xkmlm     string  `bson:"xkmlm" json:"xkmlm"` // 学科门类
	Zscc      string  `bson:"zscc" json:"zscc"`   // 招生层次 s/b/a
	SortOrder float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
}

func NewGsZydm(data GsZydmData) *GftGsZydm {
	item := &GftGsZydm{
		GsZydmData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
