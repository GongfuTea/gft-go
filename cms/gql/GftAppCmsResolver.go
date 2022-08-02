package gql

import (
	"context"
	"fmt"

	"github.com/GongfuTea/gft-go/cms/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
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

func appPosts(p graphql.ResolveParams) (any, error) {
	return mgo.CmsPostRepo.Find(context.Background(), bson.M{"state": "Published"}).All()
}

func appPost(p graphql.ResolveParams) (any, error) {
	id := p.Args["id"].(string)
	fmt.Printf("dataPost Post id, %+v", id)

	return mgo.CmsPostRepo.Get(id)
}
