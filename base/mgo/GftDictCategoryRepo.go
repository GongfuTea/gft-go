package mgo

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftDictCategoryRepo struct {
	*mgo.MgoTreeRepo
}

var DictCategoryRepo = &GftDictCategoryRepo{
	mgo.NewMgoTreeRepo("GftDictCategory", base.NewGftDictCategory),
}

// func (repo GftDictCategoryRepo) Save(model base.GftDictCategory) (*base.GftDictCategory, error) {
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	var err error

// 	fmt.Printf("%+v", model)

// 	if model.Pid != "" {
// 		parent, _ := repo.Get(model.Pid)
// 		model.Mpath = parent.Mpath + model.Slug + "."
// 	} else {
// 		model.Mpath = model.Slug + "."
// 	}

// 	if model.Id == "" {
// 		model.Id = uuid.NewString()
// 		model.CreatedAt = time.Now()
// 		_, err = repo.Coll().InsertOne(ctx, model)

// 	} else {
// 		q2 := bson.M{"$set": model}
// 		_, err = repo.Coll().UpdateByID(ctx, model.Id, q2)
// 	}

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &model, nil

// }

func (repo GftDictCategoryRepo) All() ([]base.GftDictCategory, error) {
	var results []base.GftDictCategory
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem base.GftDictCategory
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

// func (repo GftDictCategoryRepo) Get(id string) (*base.GftDictCategory, error) {
// 	var result base.GftDictCategory

// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

// 	if err := repo.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
// 		return nil, fmt.Errorf("not found")
// 	}

// 	return &result, nil
// }
