package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftCmsPostResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var CmsPostResolver = &GftCmsPostResolver{
	Query: graphql.Fields{
		"cmsPosts": &graphql.Field{
			Type:    graphql.NewList(GfCmsPostType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: dataPosts,
		},
		"cmsPost": &graphql.Field{
			Type:    GfCmsPostType,
			Args:    gql.NewArgId(),
			Resolve: dataPost,
		},
	},

	Mutation: graphql.Fields{
		"saveCmsPost": &graphql.Field{
			Type:    GfCmsPostType,
			Args:    gql.NewArgInput(GfCmsPostInput),
			Resolve: saveDataPost,
		},
		"delCmsPost": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDataPost,
		},
	},
}

func saveDataPost(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, cms.NewGftCmsPost())

	if err != nil {
		fmt.Printf("save Post err, %+v", err)
	}
	fmt.Printf("save Post, %+v", item)

	return mgo.CmsPostRepo.Save(item)
}

func dataPosts(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.CmsPostRepo.All()
}

func dataPost(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataPost Post id, %+v", id)

	return mgo.CmsPostRepo.Get(id)
}

func delDataPost(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.CmsPostRepo.Del(id)
}

var GfCmsPostType = gql.NewObjectTree("GftCmsPost", gql.FieldsConfig{
	Strings:        []string{"title", "subTitle", "note", "createdAt", "createdBy", "type", "state"},
	NonNullStrings: []string{},
	Floats:         []string{"sortOrder"},
})

var GfCmsPostInput = gql.NewInputObject("GftCmsPostInput", gql.FieldsConfig{
	Strings:        []string{"id", "note", "subTitle", "content", "type", "state"},
	NonNullStrings: []string{"title", "slug"},
	Floats:         []string{"sortOrder"},
})
