package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftAdminRepo struct {
	*mgo.MgoRepo[*GftAdmin]
}

var AdminRepo = &GftAdminRepo{
	mgo.NewMgoRepo[*GftAdmin]("GftAdmin"),
}

func (repo GftAdminRepo) Create(username string, password string) (*GftAdmin, error) {
	var admin GftAdmin

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := repo.Coll().FindOne(ctx, q).Decode(&admin); err == nil {
		return nil, fmt.Errorf("Admin already exists.")
	}

	pass, err := auth.GeneratePassword(password)
	if err != nil {
		return nil, err
	}

	admin.Id = uuid.NewString()
	t := time.Now()
	admin.CreatedAt = &t
	admin.Username = username
	admin.Password = pass

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	_, err = repo.Coll().InsertOne(ctx, admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repo GftAdminRepo) Login(username string, password string) (*auth.TokenDetails, error) {
	var err error
	var results GftAdmin

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err = repo.Coll().FindOne(ctx, q).Decode(&results); err != nil {
		return nil, fmt.Errorf("Admin not found.")
	}
	if err = auth.ComparePassword(password, results.Password); err != nil {
		return nil, fmt.Errorf("Invalid password.")
	}

	return auth.CreateToken(results.Id)
}

func (repo GftAdminRepo) Save(model GftAdmin) (*GftAdmin, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var err error

	if model.Id == "" {
		model.Id = uuid.NewString()
		t := time.Now()
		model.CreatedAt = &t
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

func (repo GftAdminRepo) All() ([]GftAdmin, error) {
	var results []GftAdmin
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem GftAdmin
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

func (repo GftAdminRepo) Del(id string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	q := bson.M{"_id": id}
	_, err := repo.Coll().DeleteOne(ctx, q)

	if err != nil {
		return false, err
	}
	return true, nil

}
