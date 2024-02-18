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
				// fmt.Printf("argType1: %+v\n", argType1.PkgPath())
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
	outputType := s.genOutput(m.Type.Out(0))

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
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		fieldType := field.Type

		if jsonTag == ",inline" {
			inlineFileds := s.genOutputFields(fieldType)
			for k, v := range inlineFileds {
				fields[k] = v
			}
			continue
		}

		// If json tag is not present, use the field name
		if jsonTag == "" {
			jsonTag = fieldName
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0] // Remove ",omitempty" if present
		}

		var graphqlType graphql.Output
		switch fieldType.Kind() {
		case reflect.String, reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			graphqlType = mapGraphqlType(fieldType)
		case reflect.Struct:
			if fieldType.Name() == "Time" {
				graphqlType = graphql.String
			} else {
				graphqlType = s.genOutputObject(fieldType)
			}
		case reflect.Slice:
			graphqlType = graphql.NewList(s.genOutput(fieldType.Elem()))
		// Add more cases as needed
		default:
			fmt.Printf("[genOutputObject] Unsupported type for field %s %s\n", fieldType.Name(), jsonTag)
			continue
		}

		fields[jsonTag] = &graphql.Field{
			Type: graphqlType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v := reflect.ValueOf(p.Source)
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				field1Val := v.FieldByName(fieldName)
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

func (s *SchemaEngine) genOutput(t reflect.Type) graphql.Output {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String, reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return mapGraphqlType(t)

	case reflect.Struct:
		return s.genOutputObject(t)

	case reflect.Slice:
		return graphql.NewList(s.genOutput(t.Elem()))

	default:
		fmt.Printf("\n\n\n\n\nUnsupported type for object %s\n\n\n\n\n\n", t.Name())
		panic("[genOutput] Unsupported type " + t.Name())
	}

}

func (s *SchemaEngine) genInputFields(t reflect.Type) graphql.InputObjectConfigFieldMap {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fields := graphql.InputObjectConfigFieldMap{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		fieldType := field.Type

		if jsonTag == ",inline" {
			inlineFileds := s.genInputFields(fieldType)
			for k, v := range inlineFileds {
				fields[k] = v
			}
			continue
		}
		if jsonTag == "" {
			jsonTag = fieldName
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0] // Remove ",omitempty" if present
		}

		var graphqlType graphql.Input
		switch fieldType.Kind() {
		case reflect.String, reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			graphqlType = mapGraphqlType(fieldType)
		case reflect.Struct:
			if fieldType.Name() == "Time" {
				graphqlType = graphql.String
			} else {
				graphqlType = s.genInputObject(fieldType)
			}
		case reflect.Slice:
			graphqlType = graphql.NewList(s.genInputObject(fieldType.Elem()))

			// Add more cases as needed

		default:
			fmt.Printf("[genInputObject] Unsupported type for field %s %s\n", fieldType.Name(), jsonTag)
			continue
		}

		fields[jsonTag] = &graphql.InputObjectFieldConfig{
			Type: graphqlType,
		}
	}

	return fields
}

func (s *SchemaEngine) genInputObject(t reflect.Type) *graphql.InputObject {
	if obj, ok := s.inputMap[t]; ok {
		return obj
	}

	s.inputMap[t] = graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   t.Name() + "Input",
		Fields: s.genInputFields(t),
	})

	return s.inputMap[t]
}

func (s *SchemaEngine) getInputFieldConfig(t reflect.Type) graphql.FieldConfigArgument {

	fmt.Printf("Type: %+v\n", t)

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

		var graphqlType graphql.Input
		switch fieldType.Kind() {
		case reflect.String, reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			graphqlType = mapGraphqlType(fieldType)

		case reflect.Struct:
			graphqlType = s.genInputObject(fieldType)
		case reflect.Slice:
			graphqlType = graphql.NewList(s.genInputObject(fieldType.Elem()))
		// Add more cases as needed
		default:
			fmt.Printf("[getInputFieldConfig] Unsupported type for field %s %s\n", fieldType.Name(), jsonTag)
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

func mapGraphqlType(t reflect.Type) *graphql.Scalar {
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
