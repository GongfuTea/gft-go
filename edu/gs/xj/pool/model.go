package pool

import (
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/xj"
)

type GftGsXjPoolItem struct {
	xj.GftGsXj `bson:",inline"`
	Xxwrq      *time.Time `bson:"xxwrq" json:"xxwrq,omitempty"` // 学信网日期
}

type GftGsXjPoolItemRepo struct {
	*mgo.MgoRepo[*GftGsXjPoolItem]
}

var GsXjPoolItemRepo = &GftGsXjPoolItemRepo{
	mgo.NewMgoRepo[*GftGsXjPoolItem]("GftGsXjPoolItem"),
}
