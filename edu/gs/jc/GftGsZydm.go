package jc

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftGsZydm struct {
	db.DbEntity `bson:",inline" json:",inline"`
	Code        string  `bson:"code" json:"code"`
	Name        string  `bson:"name" json:"name"`
	Level       int     `bson:"level" json:"level"`
	Note        string  `bson:"note" json:"note"`
	Xwlxm       string  `bson:"xwlxm" json:"xwlxm"` // 学位类型 xs/zx
	Xkmlm       string  `bson:"xkmlm" json:"xkmlm"` // 学科门类
	Zscc        string  `bson:"zscc" json:"zscc"`   // 招生层次 s/b/a
	SortOrder   float32 `bson:"sortOrder" json:"sortOrder,omitempty"`

	// Locale map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
}

func NewGsZydm(code string, name string, level int, xwlx string, zscc string) *GftGsZydm {
	it := &GftGsZydm{
		Code:  code,
		Name:  name,
		Level: level,
		Zscc:  zscc,
		Xkmlm: code[0:2],
		Xwlxm: xwlx,
	}
	it.Id = code
	it.Init()
	return it
}
