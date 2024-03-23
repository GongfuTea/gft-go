package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftCmsCategory]
}

func NewCmsCategoryRepo() *GftCmsCategoryRepo {
	return &GftCmsCategoryRepo{
		mgo.NewMgoTreeRepo[*GftCmsCategory]("GftCmsCategory"),
	}
}
