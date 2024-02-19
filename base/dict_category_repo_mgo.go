package base

import "github.com/GongfuTea/gft-go/core/mgo"

type GftDictCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftDictCategory]
}

var DictCategoryRepo = &GftDictCategoryRepo{
	mgo.NewMgoTreeRepo[*GftDictCategory]("GftDictCategory"),
}
