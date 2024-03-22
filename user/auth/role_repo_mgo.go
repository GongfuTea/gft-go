package auth

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftAuthRoleRepo struct {
	*mgo.MgoRepo[*GftAuthRole]
}

func NewAuthRoleRepo() *GftAuthRoleRepo {
	return &GftAuthRoleRepo{
		mgo.NewMgoRepo[*GftAuthRole]("GftAuthRole"),
	}
}

func (repo GftAuthRoleRepo) All() ([]GftAuthRole, error) {
	var results []GftAuthRole
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem GftAuthRole
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
