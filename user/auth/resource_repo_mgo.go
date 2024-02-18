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

var AuthResourceRepo = &GftAuthResourceRepo{
	mgo.NewMgoTreeRepo[*GftAuthResource]("GftAuthResource"),
}

// func (repo GftAuthResourceRepo) Save(model auth.GftAuthResource) (*auth.GftAuthResource, error) {
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	var err error

// 	fmt.Printf("%+v", model)

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

// func (repo GftAuthResourceRepo) Get(id string) (*auth.GftAuthResource, error) {
// 	var result auth.GftAuthResource

// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

// 	if err := repo.Coll().FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
// 		return nil, fmt.Errorf("not found")
// 	}

// 	return &result, nil
// }

// func (repo GftAuthResourceRepo) Del(id string) (bool, error) {
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

// 	q := bson.M{"_id": id}
// 	_, err := repo.Coll().DeleteOne(ctx, q)

// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil

// }
