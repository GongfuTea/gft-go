package gjjxj

import "github.com/GongfuTea/gft-go/core/db"

type GftXgJzGjjxj struct {
	db.DbEntity `bson:",inline"`
	Nd          string `bson:"nd" json:"nd,omitempty"`     // 年度
	Pycc        string `bson:"pycc" json:"pycc,omitempty"` // 层次
	Xh          string `bson:"xh" json:"xh,omitempty"`
	Xm          string `bson:"xm" json:"xm,omitempty"`
	Xb          string `bson:"xb" json:"xb,omitempty"`
	Mz          string `bson:"mz" json:"mz,omitempty"`
	Zjhm        string `bson:"zjhm" json:"zjhm,omitempty"`
	Yxs         string `bson:"yxs" json:"yxs,omitempty"`
	Zy          string `bson:"zy" json:"zy,omitempty"`
	Rxny        string `bson:"rxny" json:"rxny,omitempty"`
	Cjhjnd      string `bson:"cjhjnd" json:"cjhjnd,omitempty"` // 曾经获奖年度
	Bz          string `bson:"bz" json:"bz,omitempty"`         // 备注
	Kycg        string `bson:"kycg" json:"kycg,omitempty"`     // 科研成果
	Hjqk        string `bson:"hjqk" json:"hjqk,omitempty"`     // 获奖情况
	Qtsm        string `bson:"qtsm" json:"qtsm,omitempty"`     // 其他说明
	Sfhbps      bool   `bson:"sfhbps" json:"sfhbps,omitempty"` // 是否合并评审
	Sfhj        bool   `bson:"sfhj" json:"sfhj,omitempty"`     // 是否获奖
}
