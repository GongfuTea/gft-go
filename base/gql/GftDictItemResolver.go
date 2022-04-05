package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftDictItemResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var DictItemResolver = &GftDictItemResolver{
	Query: graphql.Fields{
		"dictItems": &graphql.Field{
			Type: graphql.NewList(GfDictItemType),
			Args: graphql.FieldConfigArgument{
				"categoryId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: DataDicts,
		},
		"dictItem": &graphql.Field{
			Type:    GfDictItemType,
			Args:    gql.NewArgId(),
			Resolve: DataDict,
		},
	},

	Mutation: graphql.Fields{
		"saveDictItem": &graphql.Field{
			Type:    GfDictItemType,
			Args:    gql.NewArgInput(GfDictItemInput),
			Resolve: saveDataDict,
		},
		"delDictItem": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDataDict,
		},
	},
}

func saveDataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(base.GftDictItem))
	if err != nil {
		fmt.Printf("save dict err, %+v", err)
	}

	return mgo.DictItemRepo.Save(item)

}

func DataDicts(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	// categoryId := p.Args["categoryId"].(string)
	return mgo.DictItemRepo.All()
}

func DataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictItemRepo.Get(id)
}

func delDataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictItemRepo.Del(id)
}

var GfDictItemType = gql.NewObjBuilder("GfDictItem").
	AddEntityFields().
	AddString("categoryId", "name", "code", "nickname", "note").
	AddInt("level").
	AddFloat("sortOrder").GetObj()

var GfDictItemInput = gql.NewInputObjBuilder("GfDictItemInput").
	AddString("id", "pid", "categoryId", "nickname", "note").
	AddNonNullString("name", "code").
	AddInt("level").
	AddFloat("sortOrder").GetObj()
