package gql

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/edu/gs/xj"
	"github.com/GongfuTea/gft-go/edu/gs/xj/mgo"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

type GsXjFilter struct {
	db.PagerFilter
	TimePoint time.Time
}

type GftGsXjResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var GsXjResolver = &GftGsXjResolver{
	Query: graphql.Fields{
		"gsXjs": &graphql.Field{
			Type:    GfGsXjFilterResp,
			Args:    gql.NewArgFilter(graphql.NewNonNull(GfGsXjFilter)),
			Resolve: gsXjs,
		},
		"gsXj": &graphql.Field{
			Type:    GfGsXjType,
			Args:    gql.NewArgId(),
			Resolve: gsXj,
		},
	},

	Mutation: graphql.Fields{
		"saveGsXj": &graphql.Field{
			Type:    GfGsXjType,
			Args:    gql.NewArgInput(GfGsXjInput),
			Resolve: saveGsXj,
		},
		"delGsXj": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delGsXj,
		},
	},
}

func saveGsXj(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(xj.GftGsXj))

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}
	fmt.Printf("save category, %+v", item)

	return mgo.GsXjRepo.Save(item)
}

func gsXjs(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	filter, _ := gql.GqlParseFilter(p, new(db.PagerFilter))
	return mgo.GsXjRepo.Find(context.Background(), bson.M{}).Page(filter)
}

func gsXj(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataCategory category id, %+v", id)

	return mgo.GsXjRepo.Get(id)
}

func delGsXj(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.GsXjRepo.Del(id)
}

var GfGsXjType = gql.NewObjBuilder("GftGsXj").
	AddEntityFields().
	AddString("xh", "xm", "sfzh", "zjlxm", "xbm", "mzm", "zzmmm", "yxsm", "zydm", "zymc", "pyccm", "xxxsm", "pyfsm", "xsdqztm", "note", "createdAt").
	AddInt("nj").
	AddFloat("xz").GetObj()

var GfGsXjInput = gql.NewInputObjBuilder("GftGsXjInput").
	AddString("id", "sfzh", "zjlxm", "xbm", "mzm", "zzmmm", "yxsm", "zydm", "zymc", "pyccm", "xxxsm", "pyfsm", "xsdqztm", "note").
	AddNonNullString("xh", "xm").
	AddInt("nj").AddFloat("xz").GetObj()

var GfGsXjFilter = gql.NewInputObjBuilder("GftGsXjFilter").
	AddInt("page", "size").AddDateTime("timePoint").GetObj()

var GfGsXjFilterResp = gql.NewObjBuilder("GftGsXjFilterResp").
	AddField(graphql.NewList(GfGsXjType), "items").
	AddInt("total").GetObj()
