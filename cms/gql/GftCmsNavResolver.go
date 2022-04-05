package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/mgo"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

type GftCmsNavResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var CmsNavResolver = &GftCmsNavResolver{
	Query: graphql.Fields{
		"cmsNavs": &graphql.Field{
			Type:    graphql.NewList(GfCmsNavType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: cmsNavs,
		},
		"CmsNav": &graphql.Field{
			Type:    GfCmsNavType,
			Args:    gql.NewArgId(),
			Resolve: cmsNav,
		},
	},

	Mutation: graphql.Fields{
		"saveCmsNav": &graphql.Field{
			Type:    GfCmsNavType,
			Args:    gql.NewArgInput(GfCmsNavInput),
			Resolve: saveCmsNav,
		},
		"delCmsNav": &graphql.Field{
			Type:    graphql.Boolean,
			Args:    gql.NewArgId(),
			Resolve: delCmsNav,
		},
	},
}

func saveCmsNav(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)

	item, err := gql.GqlParseInput(p, new(cms.GftCmsNav))

	if err != nil {
		fmt.Printf("save category err, %+v", err)
	}
	fmt.Printf("save category, %+v", item)

	return mgo.CmsNavRepo.Save(item)
}

func cmsNavs(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	return mgo.CmsNavRepo.All()
}

func cmsNav(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataCategory category id, %+v", id)

	return mgo.CmsNavRepo.Get(id)
}

func delCmsNav(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return mgo.CmsNavRepo.Del(id)
}

var GfCmsNavType = gql.NewObjBuilder("GftCmsNav").
	AddEntityTreeFields().
	AddString("name", "content", "note", "createdAt", "type", "typeId", "state", "accessLevel").
	AddField(graphql.Boolean, "newWindow").
	AddFloat("sortOrder").GetObj()

var GfCmsNavInput = gql.NewInputObjBuilder("GftCmsNavInput").
	AddString("id", "pid", "note", "content", "type", "typeId", "state", "code").
	AddNonNullString("name", "accessLevel").
	AddField(graphql.Boolean, "newWindow").
	AddFloat("sortOrder").GetObj()
