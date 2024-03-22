package auth

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftAuthResourceRepo struct {
	*mgo.MgoTreeRepo[*GftAuthResource]
}

func NewAuthResourceRepo() *GftAuthResourceRepo {
	return &GftAuthResourceRepo{
		mgo.NewMgoTreeRepo[*GftAuthResource]("GftAuthResource"),
	}
}

func (repo GftAuthResourceRepo) All() ([]GftAuthResource, error) {
	var results []GftAuthResource
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem GftAuthResource
		err = cur.Decode(&elem)
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
