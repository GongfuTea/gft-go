package jc

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftGsZydm struct {
	*db.DbEntity `bson:",inline"`
	Code         string  `bson:"code" json:"code"`
	Name         string  `bson:"name" json:"name"`
	SortOrder    float32 `bson:"sortOrder" json:"sortOrder"`
	Level        int     `bson:"level" json:"level"`
	Note         string  `bson:"note" json:"note"`
	Xwlbm        string  `bson:"xwlbm" json:"xwlbm"` // 学位类别
	Xkmlm        string  `bson:"xkmlm" json:"xkmlm"` // 学科门类

}

func NewGftGsZydm() *GftGsZydm {
	return &GftGsZydm{
		DbEntity: db.NewDbEntity(),
	}
}
