package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsImageRepo struct {
	*mgo.MgoRepo[*GftCmsImage]
}

func NewCmsImageRepo() *GftCmsImageRepo {
	return &GftCmsImageRepo{
		mgo.NewMgoRepo[*GftCmsImage]("GftCmsImage"),
	}
}
