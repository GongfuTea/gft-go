package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftCmsImageResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var CmsImageResolver = &GftCmsImageResolver{
	Query: graphql.Fields{
		"cmsImages": &graphql.Field{
			Type:    graphql.NewList(GfCmsImageType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: dataImages,
		},
		"cmsImage": &graphql.Field{
			Type:    GfCmsImageType,
			Args:    gql.NewArgId(),
			Resolve: dataImage,
		},
	},

	Mutation: graphql.Fields{
		"saveCmsImage": &graphql.Field{
			Type:    GfCmsImageType,
			Args:    gql.NewArgInput(GfCmsImageInput),
			Resolve: saveDataImage,
		},
		"delCmsImage": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delDataImage,
		},
	},
}

func saveDataImage(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(cms.GftCmsImage))

	if err != nil {
		fmt.Printf("save Image err, %+v", err)
	}
	fmt.Printf("save Image, %+v", item)

	return mgo.CmsImageRepo.Save(item)
}

func dataImages(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.CmsImageRepo.All()
}

func dataImage(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataImage Image id, %+v", id)

	return mgo.CmsImageRepo.Get(id)
}

func delDataImage(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.CmsImageRepo.Del(id)
}

var GfCmsImageType = gql.NewObjBuilder("GftCmsImage").
	AddEntityFields().
	AddString("name", "type", "url", "note", "createdAt", "createdBy").
	AddStringList("tags").
	AddDateTime("publishDate").
	AddInt("size").
	AddFloat("sortOrder").GetObj()

var GfCmsImageInput = gql.NewInputObjBuilder("GftCmsImageInput").
	AddString("id", "name", "type", "note").
	AddNonNullString("url").
	AddStringList("tags").
	AddInt("size").
	AddFloat("sortOrder").GetObj()
