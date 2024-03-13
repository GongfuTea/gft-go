package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsBannerRepo struct {
	*mgo.MgoRepo[*GftCmsBanner]
}

var CmsBannerRepo = &GftCmsBannerRepo{
	mgo.NewMgoRepo[*GftCmsBanner]("GftCmsBanner"),
}
