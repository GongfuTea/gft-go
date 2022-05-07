package gql

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {
	return gql.MergeFields(GsXjResolver.Query)
}

func MutationFields() graphql.Fields {
	return gql.MergeFields(GsXjResolver.Mutation)
}
