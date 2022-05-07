package mgo

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/user"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GftUserRepo struct {
	*mgo.MgoRepo[*user.GftUser]
}

var UserRepo = &GftUserRepo{
	mgo.NewMgoRepo[*user.GftUser]("GftUser"),
}

func (repo GftUserRepo) Create(username string, password string) (*user.GftUser, error) {
	var user user.GftUser

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := repo.Coll().FindOne(ctx, q).Decode(&user); err == nil {
		return nil, fmt.Errorf("User already exists.")
	}

	pass, err := auth.GeneratePassword(password)
	if err != nil {
		return nil, err
	}

	user.Id = uuid.NewString()
	t := time.Now()
	user.CreatedAt = &t
	user.Username = username
	user.Password = pass

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	_, err = repo.Coll().InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo GftUserRepo) Login(username string, password string) (interface{}, error) {
	var err error
	var results user.GftUser

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err = repo.Coll().FindOne(ctx, q).Decode(&results); err != nil {
		return nil, fmt.Errorf("User not found.")
	}
	if err = auth.ComparePassword(password, results.Password); err != nil {
		return nil, fmt.Errorf("Invalid password.")
	}

	return auth.CreateToken(results.Id, results.Username)
}

func (repo GftUserRepo) Save(model user.GftUser) (*user.GftUser, error) {
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

func (repo GftUserRepo) All() ([]user.GftUser, error) {
	var results []user.GftUser
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.Coll().Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem user.GftUser
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

func (repo GftUserRepo) Del(id string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	q := bson.M{"_id": id}
	_, err := repo.Coll().DeleteOne(ctx, q)

	if err != nil {
		return false, err
	}
	return true, nil

}
