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
		"dataCategories": &graphql.Field{
			Type:    graphql.NewList(GfDataCategoryType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: dataCategories,
		},
		"dataCategory": &graphql.Field{
			Type:    GfDataCategoryType,
			Args:    gql.NewArgId(),
			Resolve: dataCategory,
		},
	},

	Mutation: graphql.Fields{
		"saveDataCategory": &graphql.Field{
			Type:    GfDataCategoryType,
			Args:    gql.NewArgInput(GfDataCategoryInput),
			Resolve: saveDataCategory,
		},
		"delDataCategory": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDataCategory,
		},
	},
}

func saveDataCategory(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, base.NewGftDictCategory())

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}

	fmt.Printf("gql save category, %#v\n", item)

	return mgo.DictCategoryRepo.Save(item)
}

func dataCategories(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.DictCategoryRepo.All()
}

func dataCategory(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictCategoryRepo.Get(id)
}

func delDataCategory(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictCategoryRepo.Del(id)
}

var GfDataCategoryType = gql.NewObjectTree("GfDataCategory", gql.FieldsConfig{
	Strings:        []string{"name", "note", "createdBy"},
	NonNullStrings: []string{},
	Floats:         []string{"sortOrder"},
})

var GfDataCategoryInput = gql.NewInputObject("GfDataCategoryInput", gql.FieldsConfig{
	Strings:        []string{"id", "pid", "note"},
	NonNullStrings: []string{"name", "slug"},
	Floats:         []string{"sortOrder"},
})
