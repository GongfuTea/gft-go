package mgo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgoRepo struct {
	Name string
}

func (repo MgoRepo) Coll() *mongo.Collection {
	return DefaultMongo.Collection(repo.Name)
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
