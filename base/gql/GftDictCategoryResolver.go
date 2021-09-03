package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/core/db"
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
			Type: GfDataCategoryType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: dataCategory,
		},
	},

	Mutation: graphql.Fields{
		"saveDataCategory": &graphql.Field{
			Type: GfDataCategoryType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: GfDataCategoryInput,
				},
			},
			Resolve: saveDataCategory,
		},
		"delDataCategory": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
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

var GfDataCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GfDataCategory",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Printf("GfDataCategory source, %#v\n", p.Source)

				source, ok := p.Source.(db.IDbEntity)
				fmt.Printf("GfDataCategory ok, %#v\n", ok)
				fmt.Printf("GfDataCategory source, %#v\n", source)
				return source.ID(), nil
			},
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
		"mpath": &graphql.Field{
			Type: graphql.String,
		},
		"note": &graphql.Field{
			Type: graphql.String,
		},
		"sortOrder": &graphql.Field{
			Type: graphql.Float,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source, _ := p.Source.(db.IDbEntity)
				return source.GetCreatedAt(), nil
			},
		},
		"createdBy": &graphql.Field{
			Type: graphql.String,
		},
		// "locale": &graphql.Field{
		// 	Type: graphql.String,
		// },
	},
})

var GfDataCategoryInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GfDataCategoryInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pid": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"slug": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"note": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sortOrder": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
