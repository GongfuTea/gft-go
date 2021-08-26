package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftDictResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var DictResolver = &GftDictResolver{
	Query: graphql.Fields{
		"dataDicts": &graphql.Field{
			Type: graphql.NewList(GfDataDictType),
			Args: graphql.FieldConfigArgument{
				"categoryId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: DataDicts,
		},
		"dataDict": &graphql.Field{
			Type: GfDataDictType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: DataDict,
		},
	},

	Mutation: graphql.Fields{
		"saveDataDict": &graphql.Field{
			Type: GfDataDictType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: GfDataDictInput,
				},
			},
			Resolve: saveDataDict,
		},
		"delDataDict": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: delDataDict,
		},
	},
}

func saveDataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item := base.GftDict{}
	err := gql.GqlParseInput(p, &item)
	if err != nil {
		fmt.Printf("save dict err, %+v", err)
	}
	fmt.Printf("save dict, %+v", item)

	return mgo.DictRepo.Save(item)
}

func DataDicts(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	// categoryId := p.Args["categoryId"].(string)
	return mgo.DictRepo.All()
}

func DataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictRepo.Get(id)
}

func delDataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.DictRepo.Del(id)
}

var GfDataDictType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GfDataDict",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"categoryId": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"nickname": &graphql.Field{
			Type: graphql.String,
		},
		"level": &graphql.Field{
			Type: graphql.Int,
		},
		"note": &graphql.Field{
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

var GfDataDictInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GfDataDictInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"categoryId": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"nickname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"note": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sortOrder": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
