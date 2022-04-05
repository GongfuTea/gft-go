package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/edu/gs/xj"
	"github.com/GongfuTea/gft-go/edu/gs/xj/mgo"
	"github.com/graphql-go/graphql"
)

type GftGsXjResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var GsXjResolver = &GftGsXjResolver{
	Query: graphql.Fields{
		"gsXjs": &graphql.Field{
			Type:    graphql.NewList(GfGsXjType),
			Args:    graphql.FieldConfigArgument{},
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
	return mgo.GsXjRepo.All()
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
	AddString("xh", "xm", "note", "createdAt").
	AddFloat("xz").GetObj()

var GfGsXjInput = gql.NewInputObjBuilder("GftGsXjInput").
	AddString("id", "note").
	AddNonNullString("xh", "xm").
	AddFloat("xz").GetObj()
