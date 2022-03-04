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

func GqlParseInput(p graphql.ResolveParams, input types.IEntity) (output types.IEntity, err error) {
	args := p.Args["input"].(map[string]interface{})
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
