package gql

import (
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range GsXjResolver.Query {
		fields[k] = v
	}

	return fields
}

func MutationFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range GsXjResolver.Mutation {
		fields[k] = v
	}
	return fields
}
