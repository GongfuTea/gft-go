package gql

import (
	"context"
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/edu/gs/xj"
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

	return xj.GsXjRepo.Save(item)
}

func gsXjs(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	filter, _ := gql.GqlParseFilter(p, new(db.PagerFilter))
	return xj.GsXjRepo.Find(context.Background(), bson.M{}).Page(filter)
}

func gsXj(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataCategory category id, %+v", id)

	return xj.GsXjRepo.Get(id)
}

func delGsXj(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return xj.GsXjRepo.Del(id)
}

var gqlBuilder = gql.NewGqlBuilder(&xj.GftGsXj{})

var GfGsXjType = gqlBuilder.NewObjBuilder("GftGsXj").
	AddEntityFields().
	AddFields("xh", "xm", "nj", "xz", "sfzh", "zjlxm", "xb", "xbm", "mz", "mzm", "zzmm", "zzmmm", "yxs", "yxsm", "zydm", "zymc", "pycc", "pyccm", "xxxs", "xxxsm", "pyfs", "pyfsm", "xsdqzt", "xsdqztm", "note", "createdAt").
	Build()

var GfGsXjInput = gqlBuilder.NewInputObjBuilder("GftGsXjInput").
	AddFields("id", "nj", "xz", "sfzh", "zjlxm", "xbm", "mzm", "zzmmm", "yxsm", "zydm", "zymc", "pyccm", "xxxsm", "pyfsm", "xsdqztm", "note").
	AddNonNullFields("xh", "xm").
	Build()

var GfGsXjFilter = gql.NewInputObjBuilder("GftGsXjFilter").
	AddInt("page", "size").AddDateTime("timePoint").GetObj()

var GfGsXjFilterResp = gql.NewObjBuilder("GftGsXjFilterResp").
	AddField(graphql.NewList(GfGsXjType), "items").
	AddInt("total").GetObj()
