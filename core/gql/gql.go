package gql

import (
	"encoding/json"
	"errors"

	"github.com/graphql-go/graphql"
)

func GqlMustLogin(p graphql.ResolveParams) {
	isLogin := p.Context.Value("isLogin")
	if isLogin == nil || !isLogin.(bool) {
		panic(errors.New("Unauthorized"))
	}

}

func GqlParseInput(p graphql.ResolveParams, input interface{}) error {
	args := p.Args["input"].(map[string]interface{})
	data, _ := json.Marshal(args)
	// fmt.Printf("save category, %+v", args)

	return json.Unmarshal(data, &input)
}
