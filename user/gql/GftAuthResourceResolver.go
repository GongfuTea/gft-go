package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/mgo"
	"github.com/graphql-go/graphql"
)

type GftAuthResourceResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var AuthResourceResolver = &GftAuthResourceResolver{
	Query: graphql.Fields{
		"authResources": &graphql.Field{
			Type:    graphql.NewList(GfAuthResourceType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: authResources,
		},
		"authResource": &graphql.Field{
			Type: GfAuthResourceType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: authResource,
		},
	},

	Mutation: graphql.Fields{
		"saveAuthResource": &graphql.Field{
			Type: GfAuthResourceType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: GfAuthResourceInput,
				},
			},
			Resolve: saveAuthResource,
		},
		"delAuthResource": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: delAuthResource,
		},
	},
}

func saveAuthResource(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, auth.NewGftAuthResource())

	if err != nil {
		fmt.Printf("save resource err, %+v", err)
	}
	fmt.Printf("save resource, %+v", item)

	return mgo.AuthResourceRepo.Save(item)
}

func authResources(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.AuthResourceRepo.All()
}

func authResource(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.AuthResourceRepo.Get(id)
}

func delAuthResource(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.AuthResourceRepo.Del(id)
}

var GfAuthResourceType = gql.NewObjBuilder("GfAuthResource").
	AddEntityTreeFields().
	AddString("name", "category").
	AddField(graphql.NewList(GfAuthOperationType), "operations").
	AddFloat("sortOrder").GetObj()

var GfAuthResourceInput = gql.NewInputObjBuilder("GfAuthResourceInput").
	AddString("id", "pid", "code", "category").
	AddNonNullString("name").
	AddField(graphql.NewList(GfAuthOperationInput), "operations").
	AddFloat("sortOrder").GetObj()

var GfAuthOperationType = gql.NewObjBuilder("GfAuthOperation").
	AddNonNullString("name", "code").GetObj()

var GfAuthOperationInput = gql.NewInputObjBuilder("GfAuthOperationInput").
	AddNonNullString("name", "code").GetObj()
