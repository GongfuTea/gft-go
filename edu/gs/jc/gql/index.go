package gql

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {

	return gql.MergeFields(GsYxsResolver.Query, GsZydmResolver.Query)
}

func MutationFields() graphql.Fields {

	return gql.MergeFields(GsYxsResolver.Mutation, GsZydmResolver.Mutation)
}
