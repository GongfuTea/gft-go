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

	item := auth.GftAuthResource{}
	err := gql.GqlParseInput(p, &item)
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

var GfAuthResourceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GfAuthResource",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"pid": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"slug": &graphql.Field{
			Type: graphql.String,
		},
		"operations": &graphql.Field{
			Type: graphql.String,
		},
		"sortOrder": &graphql.Field{
			Type: graphql.Float,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var GfAuthResourceInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GfAuthResourceInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pid": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"slug": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"operations": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sortOrder": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})

var GfAuthOperationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GfAuthOperation",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"slug": &graphql.Field{
			Type: graphql.String,
		},
	},
})
