package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/mgo"
	"github.com/graphql-go/graphql"
)

type GftGsYxsResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var GsYxsResolver = &GftGsYxsResolver{
	Query: graphql.Fields{
		"gsYxses": &graphql.Field{
			Type:    graphql.NewList(GfGsYxsType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: gsYxses,
		},
		"gsYxs": &graphql.Field{
			Type:    GfGsYxsType,
			Args:    gql.NewArgId(),
			Resolve: gsYxs,
		},
	},

	Mutation: graphql.Fields{
		"saveGsYxs": &graphql.Field{
			Type:    GfGsYxsType,
			Args:    gql.NewArgInput(GfGsYxsInput),
			Resolve: saveGsYxs,
		},
		"delGsYxs": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delGsYxs,
		},
	},
}

func saveGsYxs(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(jc.GftGsYxs))

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}
	fmt.Printf("save category, %+v", item)

	return mgo.GsYxsRepo.Save(item)
}

func gsYxses(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)
	return mgo.GsYxsRepo.All()
}

func gsYxs(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.GsYxsRepo.Get(id)
}

func delGsYxs(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.GsYxsRepo.Del(id)
}

var GfGsYxsType = gql.NewObjBuilder("GftGsYxs").
	AddEntityTreeFields().
	AddString("name", "note").
	AddFloat("sortOrder").GetObj()

var GfGsYxsInput = gql.NewInputObjBuilder("GfGsYxsInput").
	AddString("id", "pid", "note").
	AddNonNullString("name", "code").
	AddFloat("sortOrder").GetObj()
