package gql

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {

	return gql.MergeFields(CmsCategoryResolver.Query, CmsPostResolver.Query, CmsImageResolver.Query, CmsNavResolver.Query, AppCmsResolver.Query)
}

func MutationFields() graphql.Fields {

	return gql.MergeFields(CmsCategoryResolver.Mutation, CmsPostResolver.Mutation, CmsImageResolver.Mutation, CmsNavResolver.Mutation, AppCmsResolver.Mutation)
}
