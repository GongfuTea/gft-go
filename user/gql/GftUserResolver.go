package gql

import (
	"github.com/GongfuTea/gft-go/user/mgo"
	"github.com/graphql-go/graphql"
)

type GftUserResolver struct {
	Query    graphql.Fields
	Mutation graphql.Fields
}

var UserResolver = &GftUserResolver{
	Query: graphql.Fields{},

	Mutation: graphql.Fields{
		"login": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				}},
			Resolve: login,
		},
	},
}

func login(p graphql.ResolveParams) (interface{}, error) {
	user := p.Args["username"].(string)
	pass := p.Args["password"].(string)
	return mgo.UserRepo.Login(user, pass)
}
