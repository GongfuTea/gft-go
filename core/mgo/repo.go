package mgo

import (
	"context"
	"fmt"
	"time"

	jsonx "github.com/GongfuTea/gft-go/core/jsonx"
	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoRepo[T types.IEntity] struct {
	Name string
}

func NewMgoRepo[T types.IEntity](name string) *MgoRepo[T] {
	return &MgoRepo[T]{Name: name}
}

func (repo MgoRepo[T]) Coll() *mongo.Collection {
	return DefaultMongo.Collection(repo.Name)
}

func (repo MgoRepo[T]) Get(id string) (T, error) {
	result := *new(T)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := repo.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(result); err != nil {
		return *new(T), fmt.Errorf("not found, %#V", err)
	}

	return result, nil
}

func (repo MgoRepo[T]) All() ([]T, error) {
	var results []T
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		elem := new(T)
		err = cur.Decode(elem)

		if err != nil {
			return nil, err
		}
		results = append(results, *elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

func (repo MgoRepo[T]) Save(model T) (T, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var err error

	jsonx.PrintAsJson("MgoRepo Save", model)
	if repo.IsExist(model.ID()) {
		q2 := bson.M{"$set": model}
		_, err = repo.Coll().UpdateByID(ctx, model.ID(), q2)
	} else {
		model.Init()
		_, err = repo.Coll().InsertOne(ctx, model)
	}

	return model, err
}

func (repo MgoRepo[T]) Del(id string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	q := bson.M{"_id": id}
	_, err := repo.Coll().DeleteOne(ctx, q)

	if err != nil {
		return false, err
	}
	return true, nil

}

func (repo MgoRepo[T]) IsExist(id string) bool {
	if id == "" {
		return false
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result := repo.Coll().FindOne(ctx, bson.M{"_id": id})

	return result.Err() == nil
}
