package mgo

import (
	"context"
	"fmt"

	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IQuery[T types.IEntity] interface {
	One() (T, error)
	All() (results []T, err error)
	Count() (int64, error)
	Exist() (bool, error)
}

type Query[T types.IEntity] struct {
	filter any
	ctx    context.Context
	coll   *mongo.Collection
}

func (q Query[T]) One() (T, error) {
	result := *new(T)

	opt := options.FindOne()

	if err := q.coll.FindOne(q.ctx, q.filter, opt).Decode(result); err != nil {
		return *new(T), fmt.Errorf("not found, %#V", err)
	}

	return result, nil
}

func (q Query[T]) All() (results []T, err error) {

	opt := options.Find()

	cur, err := q.coll.Find(q.ctx, q.filter, opt)
	if err != nil {
		return nil, err
	}

	err = cur.All(q.ctx, &results)
	return
}

func (q Query[T]) Count() (int64, error) {
	opt := options.Count()
	return q.coll.CountDocuments(q.ctx, q.filter, opt)
}

func (q Query[T]) Exist() (bool, error) {
	n, err := q.Count()
	return n > 0, err
}
