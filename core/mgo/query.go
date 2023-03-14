package mgo

import (
	"context"
	"fmt"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryPageResult[T types.IEntity] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}

type IQuery[T types.IEntity] interface {
	One() (T, error)
	All() (results []T, err error)
	Page(filter *db.PagerFilter) (result QueryPageResult[T], err error)
	Count() (int64, error)
	Exist() (bool, error)
}

type Query[T types.IEntity] struct {
	filter any
	size   *int64
	page   *int64
	ctx    context.Context
	coll   *mongo.Collection
	sort   any
}

func (q Query[T]) One() (T, error) {
	result := new(T)

	opt := options.FindOne()

	if q.sort != nil {
		opt.SetSort(q.sort)
	}

	if err := q.coll.FindOne(q.ctx, q.filter, opt).Decode(result); err != nil {
		return *new(T), fmt.Errorf("not found, %#v", err)
	}

	return *result, nil
}

func (q *Query[T]) Page(filter *db.PagerFilter) (result QueryPageResult[T], err error) {
	q.page = &filter.Page
	q.size = &filter.Size
	result.Items, err = q.All()
	result.Total, _ = q.Total()
	return
}

func (q *Query[T]) All() (results []T, err error) {

	opt := options.Find()

	if q.sort != nil {
		opt.SetSort(q.sort)
	}

	if q.page != nil && q.size != nil {
		opt.SetSkip(*q.page * *q.size)
		opt.SetLimit(*q.size)
	}

	cur, err := q.coll.Find(q.ctx, q.filter, opt)
	if err != nil {
		return nil, err
	}

	err = cur.All(q.ctx, &results)
	return
}

func (q *Query[T]) SetSort(sort any) *Query[T] {
	q.sort = sort
	return q
}

func (q *Query[T]) Count() (int64, error) {
	opt := options.Count()

	if q.page != nil && q.size != nil {
		opt.SetSkip(*q.page * *q.size)
		opt.SetLimit(*q.size)
	}

	return q.coll.CountDocuments(q.ctx, q.filter, opt)
}

func (q *Query[T]) Total() (int64, error) {
	opt := options.Count()
	return q.coll.CountDocuments(q.ctx, q.filter, opt)
}

func (q Query[T]) Exist() (bool, error) {
	n, err := q.Count()
	return n > 0, err
}
