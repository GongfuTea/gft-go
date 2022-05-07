package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type InputObjBuilder struct {
	Name   string
	Fields graphql.InputObjectConfigFieldMap
}

func NewInputObjBuilder(name string) *InputObjBuilder {
	return &InputObjBuilder{
		Name:   name,
		Fields: graphql.InputObjectConfigFieldMap{},
	}
}

func (ob *InputObjBuilder) AddString(fields ...string) *InputObjBuilder {
	return ob.AddField(graphql.String, fields...)
}

func (ob *InputObjBuilder) AddStringList(fields ...string) *InputObjBuilder {
	return ob.AddField(graphql.NewList(graphql.String), fields...)
}

func (ob *InputObjBuilder) AddNonNullString(fields ...string) *InputObjBuilder {
	return ob.AddField(graphql.NewNonNull(graphql.String), fields...)
}

func (ob *InputObjBuilder) AddFloat(fields ...string) *InputObjBuilder {
	return ob.AddField(graphql.Float, fields...)
}

func (ob *InputObjBuilder) AddInt(fields ...string) *InputObjBuilder {
	return ob.AddField(graphql.Int, fields...)
}

func (ob *InputObjBuilder) AddDateTime(fields ...string) *InputObjBuilder {
	return ob.AddField(graphql.DateTime, fields...)
}

func (ob *InputObjBuilder) AddField(t graphql.Input, fields ...string) *InputObjBuilder {
	for _, f := range fields {
		fmt.Println("input:", f)
		ob.Fields[f] = &graphql.InputObjectFieldConfig{
			Type: t,
		}
	}
	return ob
}

func (ob *InputObjBuilder) AddPageFields() *InputObjBuilder {
	return ob.AddInt("page", "size")
}

func (ob *InputObjBuilder) GetObj() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   ob.Name,
		Fields: ob.Fields,
	})

}
