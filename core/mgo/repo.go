package mgo

import (
	"context"
	"fmt"
	"time"

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
	if id == "" {
		return *new(T), fmt.Errorf("id is empty")
	}
	return repo.Find(bson.M{"_id": id}).One()
}

func (repo MgoRepo[T]) Find(filter any) IQuery[T] {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return &Query[T]{
		ctx:    ctx,
		filter: filter,
		coll:   repo.Coll(),
	}
}

func (repo MgoRepo[T]) All() ([]T, error) {
	return repo.Find(bson.M{}).All()
}

/** deprecated */
func (repo MgoRepo[T]) Save(model T) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err error

	// jsonx.PrintAsJson("MgoRepo Save", model)
	if y, _ := repo.IsExist(model.ID()); y {
		q2 := bson.M{"$set": model}
		_, err = repo.Coll().UpdateByID(ctx, model.ID(), q2)
	} else {
		model.Init()
		_, err = repo.Coll().InsertOne(ctx, model)
	}

	return model, err
}

func (repo MgoRepo[T]) Insert(model T) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := repo.Coll().InsertOne(ctx, model)
	return model, err
}

func (repo MgoRepo[T]) UpdateById(id string, m any) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := repo.Coll().UpdateByID(ctx, id, bson.M{"$set": m})
	return err == nil, err
}

func (repo MgoRepo[T]) Del(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	q := bson.M{"_id": id}
	_, err := repo.Coll().DeleteOne(ctx, q)

	if err != nil {
		return false, err
	}
	return true, nil

}

func (repo MgoRepo[T]) DelAll() (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	return repo.Coll().DeleteMany(ctx, bson.D{})
}

func (repo MgoRepo[T]) IsExist(id string) (bool, error) {
	if id == "" {
		return false, nil
	}
	return repo.Find(bson.M{"_id": id}).Exist()
}
