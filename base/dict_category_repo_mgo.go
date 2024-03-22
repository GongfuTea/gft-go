package base

import "github.com/GongfuTea/gft-go/core/mgo"

type GftDictCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftDictCategory]
}

func NewDictCategoryRepo() *GftDictCategoryRepo {
	return &GftDictCategoryRepo{
		mgo.NewMgoTreeRepo[*GftDictCategory]("GftDictCategory"),
	}
}
