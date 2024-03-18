package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsBannerRepo struct {
	*mgo.MgoRepo[*GftCmsBanner]
}

func NewBannerRepo() *GftCmsBannerRepo {
	return &GftCmsBannerRepo{
		mgo.NewMgoRepo[*GftCmsBanner]("GftCmsBanner"),
	}
}
