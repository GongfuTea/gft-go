package gql

import (
	"fmt"
	"reflect"

	"github.com/GongfuTea/gft-go/x"
	"github.com/GongfuTea/gft-go/x/tagx"
	"github.com/graphql-go/graphql"
)

func NewGqlBuilder(target any) *GqlBuilder {
	return &GqlBuilder{target: target}
}

type GqlBuilder struct {
	target any
}

func (b *GqlBuilder) NewObjBuilder(name string) *GqlObjBuilder {
	return &GqlObjBuilder{
		GqlBuilder: b,
		Ob:         NewObjBuilder(name),
		Name:       name,
	}
}

func (b *GqlBuilder) NewInputObjBuilder(name string) *GqlInputObjBuilder {
	return &GqlInputObjBuilder{
		GqlBuilder: b,
		Ob:         NewInputObjBuilder(name),
		Name:       name,
	}
}

type GqlObjBuilder struct {
	*GqlBuilder
	Ob   *ObjBuilder
	Name string
}

func (b *GqlObjBuilder) Build() *graphql.Object {

	return b.Ob.GetObj()
}

type GqlInputObjBuilder struct {
	*GqlBuilder
	Ob   *InputObjBuilder
	Name string
}

func (b *GqlInputObjBuilder) Build() *graphql.InputObject {

	return b.Ob.GetObj()
}

func (b *GqlObjBuilder) AddFields(fields ...string) *GqlObjBuilder {

	e := reflect.ValueOf(b.target).Elem()
	for i := 0; i < e.NumField(); i++ {
		name := tagx.GetJsonName(e.Type().Field(i))
		varType := e.Type().Field(i).Type

		if x.Contains(fields, name) {
			switch varType.Name() {
			case "string":
				b.Ob.AddString(name)
			case "int":
				b.Ob.AddInt(name)
			case "float64":
				b.Ob.AddFloat(name)
			default:
				panic("[BuildGqlObj] Not Supported Type: " + varType.Name())
			}
		}
		fmt.Printf("BuildGqlObj: %v %v\n", name, varType)
	}
	return b
}

func (b *GqlObjBuilder) AddNonNullFields(fields ...string) *GqlObjBuilder {

	e := reflect.ValueOf(b.target).Elem()
	for i := 0; i < e.NumField(); i++ {
		name := tagx.GetJsonName(e.Type().Field(i))
		varType := e.Type().Field(i).Type

		if x.Contains(fields, name) {
			switch varType.Name() {
			case "string":
				b.Ob.AddNonNullString(name)
			default:
				panic("[BuildGqlObj] Not Supported Type: " + varType.Name())
			}
		}
		fmt.Printf("BuildGqlObj: %v %v\n", name, varType)
	}
	return b
}

func (b *GqlObjBuilder) AddEntityFields() *GqlObjBuilder {
	b.Ob.AddEntityFields()
	return b
}

func (b *GqlInputObjBuilder) AddFields(fields ...string) *GqlInputObjBuilder {

	e := reflect.ValueOf(b.target).Elem()
	for i := 0; i < e.NumField(); i++ {
		name := tagx.GetJsonName(e.Type().Field(i))
		varType := e.Type().Field(i).Type

		if x.Contains(fields, name) {
			switch varType.Name() {
			case "string":
				b.Ob.AddString(name)
			case "int":
				b.Ob.AddInt(name)
			case "float64":
				b.Ob.AddFloat(name)
			default:
				panic("[BuildGqlInputObj] Not Supported Type: " + varType.Name())
			}
		}

		fmt.Printf("BuildGqlInputObj: %v %v\n", name, varType)
	}

	return b
}

func (b *GqlInputObjBuilder) AddNonNullFields(fields ...string) *GqlInputObjBuilder {

	e := reflect.ValueOf(b.target).Elem()
	for i := 0; i < e.NumField(); i++ {
		name := tagx.GetJsonName(e.Type().Field(i))
		varType := e.Type().Field(i).Type

		if x.Contains(fields, name) {
			switch varType.Name() {
			case "string":
				b.Ob.AddNonNullString(name)
			default:
				panic("[BuildGqlInputObj] Not Supported Type: " + varType.Name())
			}
		}
		fmt.Printf("BuildGqlInputObj: %v %v\n", name, varType)
	}
	return b
}
