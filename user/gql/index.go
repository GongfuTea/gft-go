package gql

import (
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range AuthResourceResolver.Query {
		fields[k] = v
	}

	return fields
}

func MutationFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range AuthResourceResolver.Mutation {
		fields[k] = v
	}

	return fields
}
