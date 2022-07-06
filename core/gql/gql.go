package gql

import (
	"encoding/json"
	"errors"

	"github.com/GongfuTea/gft-go/types"
	"github.com/graphql-go/graphql"
)

func GqlMustLogin(p graphql.ResolveParams) {
	isLogin := p.Context.Value("isLogin")
	if isLogin == nil || !isLogin.(bool) {
		panic(errors.New("Unauthorized"))
	}

}

func GqlParse[T any](val any, input T) (output T, err error) {
	data, _ := json.Marshal(val)
	err = json.Unmarshal(data, &input)
	return input, err
}

func GqlParseInput[T types.IEntity](p graphql.ResolveParams, input T) (output T, err error) {
	args := p.Args["input"].(map[string]any)
	data, _ := json.Marshal(args)
	// fmt.Printf("save category, %+v", args)

	err = json.Unmarshal(data, &input)
	return input, err
}

func GqlParseFilter[T any](p graphql.ResolveParams, input T) (output T, err error) {
	args := p.Args["filter"]
	data, _ := json.Marshal(args)
	// fmt.Printf("save category, %+v", args)

	err = json.Unmarshal(data, &input)
	return input, err
}

func MergeFields(list ...graphql.Fields) graphql.Fields {
	var fields = graphql.Fields{}
	for _, items := range list {
		for k, v := range items {
			fields[k] = v
		}
	}
	return fields
}
