package gql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
)

var DefaultSchemaEngine *SchemaEngine = &SchemaEngine{}

type SchemaEngine struct {
	resolvers []interface{}
}

func (s *SchemaEngine) AddResolver(resolver interface{}) {
	s.resolvers = append(s.resolvers, resolver)
}

func (s *SchemaEngine) GenerateSchema() graphql.Schema {
	var query, mutation = s.getFields()
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: query}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutation}),
	})
	return schema
}

func (s *SchemaEngine) getFields() (graphql.Fields, graphql.Fields) {
	var mutation = graphql.Fields{}
	var query = graphql.Fields{}

	for _, r := range s.resolvers {
		cmds := s.getMethods(r, "Cmd")
		for _, cmd := range cmds {
			fmt.Printf("cmd: %+v\n", cmd.Name)

			cmdName := strings.ToLower(string(cmd.Name[0])) + cmd.Name[1:]
			mutation[cmdName] = s.genFieldFromMethod(r, cmd)
		}

		queries := s.getMethods(r, "Query")
		for _, cmd := range queries {
			fmt.Printf("Query: %+v\n", cmd.Name)
			cmdName := strings.ToLower(string(cmd.Name[0])) + cmd.Name[1:]
			query[cmdName] = s.genFieldFromMethod(r, cmd)
		}
	}
	return query, mutation
}

func (s *SchemaEngine) getMethods(r any, typ string) (methods []reflect.Method) {
	t := reflect.TypeOf(r)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println("xx", method.Name, method.Type.NumIn())

		if method.Type.NumIn() == 2 { // one argument + receiver
			argType1 := method.Type.In(1) // 0 is the receiver
			if strings.HasSuffix(argType1.Name(), typ) {
				methods = append(methods, method)
			}
		}
	}
	return
}

func (s *SchemaEngine) genFieldFromMethod(resolver any, m reflect.Method) *graphql.Field {
	inputType := m.Type.In(1)
	inputArgs := s.genFieldConfig(inputType)
	outputType := s.genOutputObject(m.Type.Out(0).Elem())

	return &graphql.Field{
		Type: outputType,
		Args: inputArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			obj, err := s.convertResolveParams(inputType, p)
			if err != nil {
				return nil, err
			}

			result := m.Func.Call([]reflect.Value{reflect.ValueOf(resolver), reflect.ValueOf(obj)})
			return result[0].Interface(), nil
		},
	}
}

func (s *SchemaEngine) convertResolveParams(t reflect.Type, p graphql.ResolveParams) (any, error) {
	jsonData, err := json.Marshal(p.Args)
	if err != nil {
		return nil, err
	}

	newObj := reflect.New(t)
	err = json.Unmarshal(jsonData, newObj.Interface())
	if err != nil {
		return nil, err
	}

	return newObj.Elem().Interface(), nil
}

func (s *SchemaEngine) genOutputObject(t reflect.Type) *graphql.Object {
	fields := graphql.Fields{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		fieldType := field.Type

		// If json tag is not present, use the field name
		if jsonTag == "" {
			jsonTag = fieldName
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0] // Remove ",omitempty" if present
		}

		var graphqlType graphql.Output
		switch fieldType.Kind() {
		case reflect.String:
			graphqlType = graphql.String
		case reflect.Int:
			graphqlType = graphql.Int
		// Add more cases as needed
		default:
			continue
		}

		fields[jsonTag] = &graphql.Field{
			Type: graphqlType,
		}
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   t.Name(),
		Fields: fields,
	})
}

func (s *SchemaEngine) genInputObject(obj interface{}) *graphql.InputObject {
	t := reflect.TypeOf(obj)

	fields := graphql.InputObjectConfigFieldMap{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		fieldType := field.Type

		// If json tag is not present, use the field name
		if jsonTag == "" {
			jsonTag = fieldName
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0] // Remove ",omitempty" if present
		}

		var graphqlType graphql.Input
		switch fieldType.Kind() {
		case reflect.String:
			graphqlType = graphql.String
		case reflect.Int:
			graphqlType = graphql.Int
		case reflect.Struct:
			graphqlType = s.genInputObject(fieldType)
		// Add more cases as needed
		default:
			continue
		}

		fields[jsonTag] = &graphql.InputObjectFieldConfig{
			Type: graphqlType,
		}
	}

	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   t.Name(),
		Fields: fields,
	})
}

func (s *SchemaEngine) genFieldConfig(t reflect.Type) graphql.FieldConfigArgument {
	args := graphql.FieldConfigArgument{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		fieldType := field.Type

		// Check if the field is nullable
		isNullable := strings.Contains(jsonTag, "omitempty")
		if isNullable {
			jsonTag = strings.Split(jsonTag, ",")[0] // Remove ",omitempty"
		}

		// If json tag is not present, use the field name
		if jsonTag == "" {
			jsonTag = fieldName
		}

		var graphqlType graphql.Output
		switch fieldType.Kind() {
		case reflect.String:
			graphqlType = graphql.String
		case reflect.Int:
			graphqlType = graphql.Int
		// Add more cases as needed
		default:
			fmt.Printf("Unsupported type for field %s\n", jsonTag)
			continue
		}

		// If the field is nullable, wrap the type in graphql.NewNonNull
		if !isNullable {
			graphqlType = graphql.NewNonNull(graphqlType)
		}

		args[jsonTag] = &graphql.ArgumentConfig{
			Type: graphqlType,
		}
	}
	return args
}
