package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsNavRepo struct {
	*mgo.MgoTreeRepo[*GftCmsNav]
}

func NewCmsNavRepo() *GftCmsNavRepo {
	return &GftCmsNavRepo{
		mgo.NewMgoTreeRepo[*GftCmsNav]("GftCmsNav"),
	}
}
