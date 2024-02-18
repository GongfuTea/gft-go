package gql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
)

var DefaultSchemaEngine *SchemaEngine = &SchemaEngine{
	outputMap: make(map[reflect.Type]*graphql.Object),
	inputMap:  make(map[reflect.Type]*graphql.InputObject),
}

type SchemaEngine struct {
	resolvers []interface{}
	outputMap map[reflect.Type]*graphql.Object
	inputMap  map[reflect.Type]*graphql.InputObject
}

func (s *SchemaEngine) AddResolver(resolver interface{}) {
	s.resolvers = append(s.resolvers, resolver)
}

func (s *SchemaEngine) GenerateSchema() graphql.Schema {
	var query, mutation = s.getFields()
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: query}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: mutation}),
	})
	return schema
}

func (s *SchemaEngine) getFields() (graphql.Fields, graphql.Fields) {
	var mutation = graphql.Fields{}
	var query = graphql.Fields{}
	for _, r := range s.resolvers {
		t := reflect.TypeOf(r)
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			if method.Type.NumIn() == 2 { // one argument + receiver
				argType1 := method.Type.In(1) // 0 is the receiver
				methodName := strings.ToLower(string(method.Name[0])) + method.Name[1:]

				if strings.HasSuffix(argType1.PkgPath(), "commands") {
					mutation[methodName] = s.genFieldFromMethod(r, method)
				} else if strings.HasSuffix(argType1.PkgPath(), "queries") {
					query[methodName] = s.genFieldFromMethod(r, method)
				}
			}
		}
	}
	return query, mutation
}

func (s *SchemaEngine) genFieldFromMethod(resolver any, m reflect.Method) *graphql.Field {
	inputType := m.Type.In(1)
	inputArgs := s.getInputFieldConfig(inputType)
	outputType := s.mapOutputType(m.Type.Out(0))

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

func (s *SchemaEngine) genOutputFields(t reflect.Type) graphql.Fields {
	fields := graphql.Fields{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName, fieldType, isInline, isNullable := s.getFieldInfo(field)

		if isInline {
			inlineFileds := s.genOutputFields(fieldType)
			for k, v := range inlineFileds {
				fields[k] = v
			}
			continue
		}

		var graphqlType graphql.Output = s.mapOutputType(fieldType)
		if graphqlType == nil {
			continue
		}

		if !isNullable {
			graphqlType = graphql.NewNonNull(graphqlType)
		}

		fields[fieldName] = &graphql.Field{
			Type: graphqlType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v := reflect.ValueOf(p.Source)
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				field1Val := v.FieldByName(field.Name)
				if field1Val.IsValid() {
					return field1Val.Interface(), nil
				}
				return nil, nil
			},
		}
	}
	return fields
}

func (s *SchemaEngine) genOutputObject(t reflect.Type) *graphql.Object {
	if obj, ok := s.outputMap[t]; ok {
		return obj
	}

	s.outputMap[t] = graphql.NewObject(graphql.ObjectConfig{
		Name:   t.Name(),
		Fields: s.genOutputFields(t),
	})

	return s.outputMap[t]
}

func (s *SchemaEngine) mapOutputType(t reflect.Type) graphql.Output {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String, reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return s.mapGraphqlType(t)

	case reflect.Struct:
		if t.Name() == "Time" {
			return graphql.String
		} else {
			return s.genOutputObject(t)
		}

	case reflect.Slice:
		return graphql.NewList(s.mapOutputType(t.Elem()))

	default:
		fmt.Printf("[mapOutputType] Unsupported type for object %s\n", t.Name())
		return nil
	}
}

func (s *SchemaEngine) mapInputType(t reflect.Type) graphql.Input {
	switch t.Kind() {
	case reflect.String, reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return s.mapGraphqlType(t)

	case reflect.Struct:
		return s.genInputObject(t)

	case reflect.Slice:
		return graphql.NewList(s.mapInputType(t.Elem()))

	default:
		fmt.Printf("[mapInputType] Unsupported type for object %s\n", t.Name())
		return nil
	}
}

func (s *SchemaEngine) getFieldInfo(field reflect.StructField) (string, reflect.Type, bool, bool) {
	fieldName := field.Tag.Get("json")
	fieldType := field.Type
	isInline := fieldName == ",inline"
	isNullable := strings.Contains(fieldName, "omitempty")
	if isNullable {
		fieldName = strings.Split(fieldName, ",")[0] // Remove ",omitempty"
	}
	if fieldName == "" {
		fieldName = field.Name
	}
	return fieldName, fieldType, isInline, isNullable
}

func (s *SchemaEngine) genInputFields(t reflect.Type) graphql.InputObjectConfigFieldMap {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fields := graphql.InputObjectConfigFieldMap{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName, fieldType, isInline, isNullable := s.getFieldInfo(field)

		if isInline {
			inlineFileds := s.genInputFields(fieldType)
			for k, v := range inlineFileds {
				fields[k] = v
			}
			continue
		}

		var graphqlType graphql.Input = s.mapInputType(fieldType)
		if graphqlType == nil {
			continue
		}

		// If the field is nullable, wrap the type in graphql.NewNonNull
		if !isNullable {
			graphqlType = graphql.NewNonNull(graphqlType)
		}

		fields[fieldName] = &graphql.InputObjectFieldConfig{
			Type: graphqlType,
		}
	}

	return fields
}

func (s *SchemaEngine) genInputObject(t reflect.Type) *graphql.InputObject {
	if obj, ok := s.inputMap[t]; ok {
		return obj
	}
	name := t.Name()
	if !strings.HasSuffix(name, "Input") {
		name = name + "Input"
	}
	s.inputMap[t] = graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   name,
		Fields: s.genInputFields(t),
	})

	return s.inputMap[t]
}

func (s *SchemaEngine) getInputFieldConfig(t reflect.Type) graphql.FieldConfigArgument {
	fields := s.genInputFields(t)
	args := graphql.FieldConfigArgument{}
	for k, v := range fields {
		args[k] = &graphql.ArgumentConfig{
			Type: v.Type,
		}
	}
	return args
}

func (s *SchemaEngine) mapGraphqlType(t reflect.Type) *graphql.Scalar {
	switch t.Kind() {
	case reflect.String:
		return graphql.String
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return graphql.Int

	case reflect.Float32, reflect.Float64:
		return graphql.Float

	case reflect.Bool:
		return graphql.Boolean

	// Add more cases as needed
	default:
		fmt.Printf("[mapGraphqlType] Unsupported type for field %s\n", t.Name())
		return nil
	}
}
