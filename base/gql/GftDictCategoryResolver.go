package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftDictCategoryResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var DictCategoryResolver = &GftDictCategoryResolver{
	Query: graphql.Fields{
		"dictCategories": &graphql.Field{
			Type:    graphql.NewList(GfDictCategoryType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: dictCategories,
		},
		"dictCategory": &graphql.Field{
			Type:    GfDictCategoryType,
			Args:    gql.NewArgId(),
			Resolve: dictCategory,
		},
	},

	Mutation: graphql.Fields{
		"saveDictCategory": &graphql.Field{
			Type:    GfDictCategoryType,
			Args:    gql.NewArgInput(GfDictCategoryInput),
			Resolve: saveDictCategory,
		},
		"delDictCategory": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDictCategory,
		},
	},
}

func saveDictCategory(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(base.GftDictCategory))

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}

	fmt.Printf("gql save category, %#v\n", item)

	return mgo.DictCategoryRepo.Save(item)
}

func dictCategories(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.DictCategoryRepo.All()
}

func dictCategory(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictCategoryRepo.Get(id)
}

func delDictCategory(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictCategoryRepo.Del(id)
}

var GfDictCategoryType = gql.NewObjBuilder("GfDictCategory").
	AddEntityTreeFields().
	AddString("name", "note").
	AddFloat("sortOrder").GetObj()

var GfDictCategoryInput = gql.NewInputObjBuilder("GfDictCategoryInput").
	AddString("id", "pid", "note").
	AddNonNullString("name", "code").
	AddFloat("sortOrder").GetObj()
