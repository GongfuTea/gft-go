package gql

import (
	"github.com/graphql-go/graphql"
)

func QueryFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range CmsCategoryResolver.Query {
		fields[k] = v
	}
	// for k, v := range DictResolver.Query {
	// 	fields[k] = v
	// }
	return fields
}

func MutationFields() graphql.Fields {
	var fields = graphql.Fields{}
	for k, v := range CmsCategoryResolver.Mutation {
		fields[k] = v
	}
	// for k, v := range DictResolver.Mutation {
	// 	fields[k] = v
	// }
	return fields
}
