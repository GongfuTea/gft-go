package mgo

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/jsonx"
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftDictRepo struct {
	*mgo.MgoRepo
}

var DictRepo = &GftDictRepo{
	mgo.NewMgoRepo("GftDict", base.NewGftDict),
}

func (repo GftDictRepo) All() ([]types.IEntity, error) {
	var results []types.IEntity
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		elem := base.NewGftDict()
		err = cur.Decode(elem)

		jsonx.PrintAsJson(elem)

		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}
