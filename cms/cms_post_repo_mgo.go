package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsPostRepo struct {
	*mgo.MgoRepo[*GftCmsPost]
}
 
func NewCmsPostRepo() *GftCmsPostRepo {
	return &GftCmsPostRepo{
		mgo.NewMgoRepo[*GftCmsPost]("GftCmsPost"),
	}
}
