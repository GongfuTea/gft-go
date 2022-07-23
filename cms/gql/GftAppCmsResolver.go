package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/cms/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftAppCmsResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var AppCmsResolver = &GftAppCmsResolver{
	Query: graphql.Fields{
		"appCmsPosts": &graphql.Field{
			Type:    graphql.NewList(GfCmsPostType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: appPosts,
		},
		"appCmsPost": &graphql.Field{
			Type:    GfCmsPostType,
			Args:    gql.NewArgId(),
			Resolve: appPost,
		},
	},
}

func appPosts(p graphql.ResolveParams) (interface{}, error) {
	return mgo.CmsPostRepo.All()
}

func appPost(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	fmt.Printf("dataPost Post id, %+v", id)

	return mgo.CmsPostRepo.Get(id)
}
