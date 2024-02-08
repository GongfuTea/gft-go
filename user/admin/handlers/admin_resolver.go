package handlers

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/user/admin"
	"github.com/graphql-go/graphql"
)

type AdminResolver struct {
}

func (r *AdminResolver) CreateAdmin() {
}

func (r *AdminResolver) UpdateAdmin() {
}

func (r *AdminResolver) FindAdmins() {
}

func (r *AdminResolver) QueryFields() graphql.Fields {
	var fields = graphql.Fields{}

	// for k, v := range UserResolver.Query {
	// 	fields[k] = v
	// }
	return fields
}

func (r *AdminResolver) MutationFields() graphql.Fields {
	var fields = graphql.Fields{
		"adminLogin": &graphql.Field{
			Type: GfAuthToken,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				}},
			Resolve: r.login,
		},
	}

	return fields
}

func (r *AdminResolver) login(p graphql.ResolveParams) (interface{}, error) {
	println("sss")
	user := p.Args["username"].(string)
	pass := p.Args["password"].(string)
	return admin.AdminRepo.Login(user, pass)
}

var GfAuthToken = gql.NewObjBuilder("GfAuthToken2").
	AddString("accessToken", "refreshToken").GetObj()
