package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftCmsCategoryResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var CmsCategoryResolver = &GftCmsCategoryResolver{
	Query: graphql.Fields{
		"cmsCategories": &graphql.Field{
			Type:    graphql.NewList(GfCmsCategoryType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: dataCategories,
		},
		"cmsCategory": &graphql.Field{
			Type:    GfCmsCategoryType,
			Args:    gql.NewArgId(),
			Resolve: dataCategory,
		},
	},

	Mutation: graphql.Fields{
		"saveCmsCategory": &graphql.Field{
			Type:    GfCmsCategoryType,
			Args:    gql.NewArgInput(GfCmsCategoryInput),
			Resolve: saveDataCategory,
		},
		"delCmsCategory": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDataCategory,
		},
	},
}

func saveDataCategory(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(cms.GftCmsCategory))

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}
	fmt.Printf("save category, %+v", item)

	return mgo.CmsCategoryRepo.Save(item)
}

func dataCategories(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)
	return mgo.CmsCategoryRepo.All()
}

func dataCategory(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataCategory category id, %+v", id)

	return mgo.CmsCategoryRepo.Get(id)
}

func delDataCategory(p graphql.ResolveParams) (any, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.CmsCategoryRepo.Del(id)
}

var GfCmsCategoryType = gql.NewObjBuilder("GftCmsCategory").
	AddEntityTreeFields().
	AddString("name", "note", "createdBy").
	AddFloat("sortOrder").GetObj()

var GfCmsCategoryInput = gql.NewInputObjBuilder("GftCmsCategoryInput").
	AddString("id", "pid", "note").
	AddNonNullString("name", "code").
	AddFloat("sortOrder").GetObj()
