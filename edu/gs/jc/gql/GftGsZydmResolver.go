package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/mgo"
	"github.com/graphql-go/graphql"
)

type GftGsZydmResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var GsZydmResolver = &GftGsZydmResolver{
	Query: graphql.Fields{
		"gsZydms": &graphql.Field{
			Type:    graphql.NewList(GfGsZydmType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: gsZydms,
		},
		"gsZydm": &graphql.Field{
			Type:    GfGsZydmType,
			Args:    gql.NewArgId(),
			Resolve: gsZydm,
		},
	},

	Mutation: graphql.Fields{
		"saveGsZydm": &graphql.Field{
			Type:    GfGsZydmType,
			Args:    gql.NewArgInput(GfGsZydmInput),
			Resolve: saveGsZydm,
		},
		"delGsZydm": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delGsZydm,
		},
	},
}

func saveGsZydm(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(jc.GftGsZydm))

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}
	fmt.Printf("save category, %+v", item)

	return mgo.GsZydmRepo.Save(item)
}

func gsZydms(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.GsZydmRepo.All()
}

func gsZydm(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.GsZydmRepo.Get(id)
}

func delGsZydm(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.GsZydmRepo.Del(id)
}

var GfGsZydmType = gql.NewObjBuilder("GftGsZydm").
	AddEntityFields().
	AddString("name", "code", "xwlxm", "xkmlm", "note").
	AddInt("level").
	AddFloat("sortOrder").GetObj()

var GfGsZydmInput = gql.NewInputObjBuilder("GfGsZydmInput").
	AddString("id", "pid", "note", "xwlxm", "xkmlm").
	AddNonNullString("name", "code").
	AddInt("level").
	AddFloat("sortOrder").GetObj()
