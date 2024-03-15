package xj

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

type GftGsXjRepo struct {
	*mgo.MgoRepo[*GftGsXj]
}

var GsXjRepo = &GftGsXjRepo{
	mgo.NewMgoRepo[*GftGsXj]("GftGsXj"),
}

func (p *GftGsXjRepo) FindZxs() (GftGsXjList, error) {
	return p.Find(bson.M{"tlEnd": nil}).All()
}

func (p *GftGsXjRepo) FindByXh(xh string) (GftGsXjList, error) {
	return p.Find(bson.M{"xh": xh}).All()
}

func (p *GftGsXjRepo) FindByXhLast(xh string) (*GftGsXj, error) {
	ls, err := p.Find(bson.M{"xh": xh}).All()
	if err != nil {
		return nil, err
	}

	return ls[0], nil
}
