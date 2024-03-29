package gql

import (
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/graphql-go/graphql"
)

type FieldsConfig struct {
	Strings        []string
	NonNullStrings []string
	Floats         []string
	Ints           []string
}

func NewArgId() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}
}

func NewArgInput(t graphql.Input) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: t,
		},
	}
}

func NewArgFilter(t graphql.Input) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"filter": &graphql.ArgumentConfig{
			Type: t,
		},
	}
}

func NewObject(name string, fields FieldsConfig) *graphql.Object {

	m := graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source, _ := p.Source.(db.IDbEntity)
				return source.ID(), nil
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source, _ := p.Source.(db.IDbEntity)
				return source.GetCreatedAt(), nil
			},
		},
	}

	for _, f := range fields.Strings {
		m[f] = &graphql.Field{
			Type: graphql.String,
		}
	}
	for _, f := range fields.NonNullStrings {
		m[f] = &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		}
	}
	for _, f := range fields.Floats {
		m[f] = &graphql.Field{
			Type: graphql.Float,
		}
	}
	for _, f := range fields.Ints {
		m[f] = &graphql.Field{
			Type: graphql.Int,
		}
	}
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: m,
	})
}

func NewObjectTree(name string, fields FieldsConfig) *graphql.Object {

	obj := NewObject(name, fields)

	obj.AddFieldConfig("pid", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbTreeEntity)
			return source.PID(), nil
		},
	})

	obj.AddFieldConfig("code", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbTreeEntity)
			return source.GetCode(), nil
		},
	})

	obj.AddFieldConfig("mpath", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbTreeEntity)
			return source.GetMpath(), nil
		},
	})
	return obj

}

func NewInputObject(name string, fields FieldsConfig) *graphql.InputObject {

	m := graphql.InputObjectConfigFieldMap{}

	for _, f := range fields.Strings {
		m[f] = &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		}
	}
	for _, f := range fields.NonNullStrings {
		m[f] = &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		}
	}
	for _, f := range fields.Floats {
		m[f] = &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		}
	}

	for _, f := range fields.Ints {
		m[f] = &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		}
	}
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   name,
		Fields: m,
	})
}
