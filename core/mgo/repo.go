package mgo

import (
	"context"
	"time"

	jsonx "github.com/GongfuTea/gft-go/core/jsonx"
	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	return repo.Find(context.Background(), bson.M{"_id": id}).One()
}

func (repo MgoRepo[T]) Find(ctx context.Context, filter any) IQuery[T] {
	println("repo find", filter)
	return &Query[T]{
		ctx:    ctx,
		filter: filter,
		coll:   repo.Coll(),
	}
}

func (repo MgoRepo[T]) All() ([]T, error) {
	return repo.Find(context.Background(), bson.M{}).All()
}

func (repo MgoRepo[T]) Save(model T) (T, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var err error

	jsonx.PrintAsJson("MgoRepo Save", model)
	if y, _ := repo.IsExist(model.ID()); y {
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

func (repo MgoRepo[T]) IsExist(id string) (bool, error) {
	if id == "" {
		return false, nil
	}
	return repo.Find(context.Background(), bson.M{"_id": id}).Exist()
}
