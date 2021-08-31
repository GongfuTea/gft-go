package mgo

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftCmsCategoryRepo struct {
	*mgo.MgoRepo
}

var CmsCategoryRepo = &GftCmsCategoryRepo{
	&mgo.MgoRepo{
		Name: "GftDictCategory",
	},
}

func (repo GftCmsCategoryRepo) Save(model cms.GftCmsCategory) (*cms.GftCmsCategory, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var err error

	fmt.Printf("%+v", model)

	if model.Pid != "" {
		parent, _ := repo.Get(model.Pid)
		model.Mpath = parent.Mpath + model.Slug + "."
	} else {
		model.Mpath = model.Slug + "."
	}

	if model.Id == "" {
		model.Id = uuid.NewString()
		model.CreatedAt = time.Now()
		_, err = repo.Coll().InsertOne(ctx, model)

	} else {
		q2 := bson.M{"$set": model}
		_, err = repo.Coll().UpdateByID(ctx, model.Id, q2)
	}

	if err != nil {
		return nil, err
	}
	return &model, nil

}

func (repo GftCmsCategoryRepo) All() ([]cms.GftCmsCategory, error) {
	var results []cms.GftCmsCategory
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem cms.GftCmsCategory
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

func (repo GftCmsCategoryRepo) Get(id string) (*cms.GftCmsCategory, error) {
	var result cms.GftCmsCategory

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := repo.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, fmt.Errorf("not found")
	}

	return &result, nil
}
