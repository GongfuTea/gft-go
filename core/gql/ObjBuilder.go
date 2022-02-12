package gql

import (
	"fmt"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/graphql-go/graphql"
)

type ObjBuilder struct {
	Obj *graphql.Object
}

func NewObjBuilder(name string) *ObjBuilder {
	cfg := graphql.ObjectConfig{
		Name:   name,
		Fields: graphql.Fields{},
	}
	return &ObjBuilder{
		Obj: graphql.NewObject(cfg),
	}
}

func (ob *ObjBuilder) AddString(fields ...string) *ObjBuilder {
	return ob.AddField(graphql.String, fields...)
}

func (ob *ObjBuilder) AddStringList(fields ...string) *ObjBuilder {
	return ob.AddField(graphql.NewList(graphql.String), fields...)
}

func (ob *ObjBuilder) AddNonNullString(fields ...string) *ObjBuilder {
	return ob.AddField(graphql.NewNonNull(graphql.String), fields...)
}

func (ob *ObjBuilder) AddFloat(fields ...string) *ObjBuilder {
	return ob.AddField(graphql.Float, fields...)
}

func (ob *ObjBuilder) AddInt(fields ...string) *ObjBuilder {
	return ob.AddField(graphql.Int, fields...)
}

func (ob *ObjBuilder) AddDateTime(fields ...string) *ObjBuilder {
	return ob.AddField(graphql.DateTime, fields...)
}

func (ob *ObjBuilder) AddField(t graphql.Output, fields ...string) *ObjBuilder {
	for _, f := range fields {
		fmt.Println("ob:", f)

		ob.Obj.AddFieldConfig(f, &graphql.Field{
			Type: t,
		})
	}
	return ob
}

func (ob *ObjBuilder) AddEntityFields() *ObjBuilder {
	ob.Obj.AddFieldConfig("id", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbEntity)
			return source.ID(), nil
		},
	})

	ob.Obj.AddFieldConfig("createdAt", &graphql.Field{
		Type: graphql.DateTime,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbEntity)
			return source.GetCreatedAt(), nil
		},
	})

	return ob
}

func (ob *ObjBuilder) AddEntityTreeFields() *ObjBuilder {

	ob.Obj.AddFieldConfig("pid", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbTreeEntity)
			if source == nil {
				return nil, nil
			}
			return source.PID(), nil
		},
	})

	ob.Obj.AddFieldConfig("code", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Printf("tree code, %+v\n", p.Source)

			source, _ := p.Source.(db.IDbTreeEntity)
			if source == nil {
				return nil, nil
			}
			return source.GetCode(), nil
		},
	})

	ob.Obj.AddFieldConfig("mpath", &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			source, _ := p.Source.(db.IDbTreeEntity)
			return source.GetMpath(), nil
		},
	})

	return ob
}

func (ob *ObjBuilder) GetObj() *graphql.Object {
	fmt.Printf("cms post, %+v\n", ob.Obj)

	return ob.Obj
}
