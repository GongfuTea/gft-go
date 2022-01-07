package gql

import (
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range CmsCategoryResolver.Query {
		fields[k] = v
	}
	for k, v := range CmsPostResolver.Query {
		fields[k] = v
	}
	for k, v := range CmsNavResolver.Query {
		fields[k] = v
	}
	return fields
}

func MutationFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range CmsCategoryResolver.Mutation {
		fields[k] = v
	}
	for k, v := range CmsPostResolver.Mutation {
		fields[k] = v
	}
	for k, v := range CmsNavResolver.Mutation {
		fields[k] = v
	}
	return fields
}
