package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type GftAdminRepo struct {
	*mgo.MgoRepo[*GftAuthAdmin]
}

func NewAdminRepo() *GftAdminRepo {
	return &GftAdminRepo{
		mgo.NewMgoRepo[*GftAuthAdmin]("GftAdmin"),
	}
}

func (repo GftAdminRepo) Create(username string, password string) (*GftAuthAdmin, error) {
	var admin GftAuthAdmin

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
	admin.CreatedAt = time.Now()
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
	var results GftAuthAdmin

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
