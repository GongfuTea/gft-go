package user

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type GftUserRepo struct {
	*mgo.MgoRepo[*GftUser]
}

func NewUserRepo() *GftUserRepo {
	return &GftUserRepo{
		mgo.NewMgoRepo[*GftUser]("GftUser"),
	}
}

func (repo GftUserRepo) Create(username string, password string) (*GftUser, error) {
	var user GftUser

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
	user.CreatedAt = time.Now()
	user.Username = username
	user.Password = pass

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	_, err = repo.Coll().InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo GftUserRepo) Login(username string, password string) (*auth.TokenDetails, error) {
	var err error
	var results GftUser

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err = repo.Coll().FindOne(ctx, q).Decode(&results); err != nil {
		return nil, fmt.Errorf("User not found.")
	}
	if err = auth.ComparePassword(password, results.Password); err != nil {
		return nil, fmt.Errorf("Invalid password.")
	}

	return auth.CreateToken(results.Id)
}
