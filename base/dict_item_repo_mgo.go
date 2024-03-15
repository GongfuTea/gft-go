package base

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

type GftDictItemRepo struct {
	*mgo.MgoRepo[*GftDictItem]
}

var DictItemRepo = &GftDictItemRepo{
	mgo.NewMgoRepo[*GftDictItem]("GftDictItem"),
}

func (r *GftDictItemRepo) FindByCategoryId(categoryId string) ([]*GftDictItem, error) {
	return r.Find(bson.M{"categoryId": categoryId}).All()
}

func (r *GftDictItemRepo) FindByItemCode(categoryId string, itemCode string) (*GftDictItem, error) {
	return r.Find(bson.M{"categoryId": categoryId, "code": itemCode}).One()
}

func (r *GftDictItemRepo) FindByItemName(categoryId string, itemValue string) (*GftDictItem, error) {
	return r.Find(bson.M{"categoryId": categoryId, "name": itemValue}).One()
}
