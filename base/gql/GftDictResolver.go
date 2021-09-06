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
			Type:    GfDataDictType,
			Args:    gql.NewArgId(),
			Resolve: DataDict,
		},
	},

	Mutation: graphql.Fields{
		"saveDataDict": &graphql.Field{
			Type:    GfDataDictType,
			Args:    gql.NewArgInput(GfDataDictInput),
			Resolve: saveDataDict,
		},
		"delDataDict": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDataDict,
		},
	},
}

func saveDataDict(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, base.NewGftDict())
	if err != nil {
		fmt.Printf("save dict err, %+v", err)
	}

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

var GfDataDictType = gql.NewObject("GfDataDict", gql.FieldsConfig{
	Strings:        []string{"categoryId", "name", "code", "nickname", "note"},
	NonNullStrings: []string{},
	Floats:         []string{"sortOrder"},
	Ints:           []string{"level"},
})

var GfDataDictInput = gql.NewInputObject("GfDataDictInput", gql.FieldsConfig{
	Strings:        []string{"id", "categoryId", "nickname", "note"},
	NonNullStrings: []string{"name", "code"},
	Floats:         []string{"sortOrder"},
})
