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

type MgoRepo struct {
	Name    string
	factory func() types.IEntity
}

func NewMgoRepo(name string, factory func() types.IEntity) *MgoRepo {
	return &MgoRepo{Name: name, factory: factory}
}

func (repo MgoRepo) Coll() *mongo.Collection {
	return DefaultMongo.Collection(repo.Name)
}

func (repo MgoRepo) Get(id string) (types.IEntity, error) {
	result := repo.factory()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := repo.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(result); err != nil {
		return nil, fmt.Errorf("not found, %#V", err)
	}

	return result, nil
}

func (repo MgoRepo) All() ([]types.IEntity, error) {
	var results []types.IEntity
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		elem := repo.factory()
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

func (repo MgoRepo) Save(model types.IEntity) (types.IEntity, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var err error
	fmt.Printf("MgoRepo Save, %#v\n", model)

	jsonx.PrintAsJson(model)
	if model.IsNew() {
		model.Init()
		jsonx.PrintAsJson(model)
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
