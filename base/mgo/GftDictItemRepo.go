package mgo

import (
	"context"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

type GftDictItemRepo struct {
	*mgo.MgoRepo[*base.GftDictItem]
}

var DictItemRepo = &GftDictItemRepo{
	mgo.NewMgoRepo[*base.GftDictItem]("GftDictItem"),
}

func (r *GftDictItemRepo) FindByCategoryId(categoryId string) ([]*base.GftDictItem, error) {
	return r.Find(context.Background(), bson.M{"categoryId": categoryId}).All()
}

func (r *GftDictItemRepo) FindByItemCode(categoryId string, itemCode string) (*base.GftDictItem, error) {
	return r.Find(context.Background(), bson.M{"categoryId": categoryId, "code": itemCode}).One()
}

func (r *GftDictItemRepo) FindByItemName(categoryId string, itemValue string) (*base.GftDictItem, error) {
	return r.Find(context.Background(), bson.M{"categoryId": categoryId, "name": itemValue}).One()
}
