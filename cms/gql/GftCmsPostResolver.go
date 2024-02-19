package gql

import (
	"context"
	"fmt"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

type GftCmsPostFilter struct {
	db.PagerFilter
	Category *string
}

type GftCmsPostResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var CmsPostResolver = &GftCmsPostResolver{
	Query: graphql.Fields{
		"cmsPosts": &graphql.Field{
			Type:    GfCmsPostFilterResp,
			Args:    gql.NewArgFilter(graphql.NewNonNull(GfCmsPostFilter)),
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

	item, err := gql.GqlParseInput(p, new(cms.GftCmsPost))

	if err != nil {
		fmt.Printf("save Post err, %+v", err)
	}
	fmt.Printf("save Post, %+v", item)

	return cms.CmsPostRepo.Save(item)
}

func dataPosts(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	filter, _ := gql.GqlParseFilter(p, new(GftCmsPostFilter))
	if filter.Category == nil {
		return cms.CmsPostRepo.Find(context.Background(), bson.M{}).Page(&filter.PagerFilter)
	} else if *filter.Category == "" {
		return cms.CmsPostRepo.Find(context.Background(), bson.M{"categoryIds": []string{}}).Page(&filter.PagerFilter)
	} else {
		return cms.CmsPostRepo.Find(context.Background(), bson.M{"categoryIds": *filter.Category}).Page(&filter.PagerFilter)
	}
}

func dataPost(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	fmt.Printf("dataPost Post id, %+v", id)

	return cms.CmsPostRepo.Get(id)
}

func delDataPost(p graphql.ResolveParams) (interface{}, error) {
	gql.GqlMustLogin(p)
	id := p.Args["id"].(string)
	return cms.CmsPostRepo.Del(id)
}

var GfCmsPostType = gql.NewObjBuilder("GftCmsPost").
	AddEntityFields().
	AddString("title", "subTitle", "content", "note", "createdAt", "createdBy", "type", "state", "format", "abstract", "accessLevel").
	AddStringList("categoryIds", "tags", "publishDepts").
	AddDateTime("publishDate").
	AddField(graphql.Boolean, "newWindow").
	AddFloat("sortOrder").GetObj()

var GfCmsPostInput = gql.NewInputObjBuilder("GftCmsPostInput").
	AddString("id", "note", "subTitle", "content", "type", "state", "format", "slug", "abstract").
	AddNonNullString("title", "accessLevel").
	AddStringList("categoryIds", "tags", "publishDepts").
	AddDateTime("publishDate").
	AddField(graphql.Boolean, "newWindow").
	AddFloat("sortOrder").GetObj()

var GfCmsPostFilter = gql.NewInputObjBuilder("GfCmsPostFilter").
	AddInt("page", "size").AddString("category").GetObj()

var GfCmsPostFilterResp = gql.NewObjBuilder("GfCmsPostFilterResp").
	AddField(graphql.NewList(GfCmsPostType), "items").
	AddInt("total").GetObj()
