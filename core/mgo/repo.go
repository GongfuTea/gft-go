package mgo

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgoRepo struct {
	Name string
}

func (repo MgoRepo) Coll() *mongo.Collection {
	return DefaultMongo.Collection(repo.Name)
}

func (repo MgoRepo) Save(model types.IEntity) (types.IEntity, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var err error

	if model.IsNew() {
		model.Init()
		_, err = repo.Coll().InsertOne(ctx, model)
	} else {
		q2 := bson.M{"$set": model}
		_, err = repo.Coll().UpdateByID(ctx, model.ID(), q2)
	}

	return model, err
}

func (repo MgoRepo) Del(id string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	q := bson.M{"_id": id}
	_, err := repo.Coll().DeleteOne(ctx, q)

	if err != nil {
		return false, err
	}
	return true, nil

}
