package gql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var AnyValueType = graphql.NewScalar(graphql.ScalarConfig{
	Name: "AnyValue",
	// Serialize will be used to convert the Golang value into the value sent to the client
	Serialize: func(value interface{}) interface{} {
		return value
	},
	// ParseValue will be used to convert the value received from the client into the Golang value
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	// ParseLiteral will be used to convert the AST literal received from the client into the Golang value
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		case *ast.IntValue:
			intVal, _ := strconv.ParseInt(valueAST.Value, 10, 64)
			return intVal
		case *ast.FloatValue:
			floatVal, _ := strconv.ParseFloat(valueAST.Value, 64)
			return floatVal
		case *ast.BooleanValue:
			return valueAST.Value
		default:
			return nil
		}
	},
})

var MapType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MapType",
	Fields: graphql.Fields{
		"key": &graphql.Field{
			Type: graphql.String,
		},
		"value": &graphql.Field{
			Type: AnyValueType,
		},
	},
})

var TimeType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "TimeType",
	Description: "The `Time` scalar type represents time in ISO 8601 format.",
	// Serialize receives a Time and converts it to a string
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case time.Time:
			return value.Format("2006-01-02T15:04:05.000Z")
		default:
			return nil
		}
	},
	// ParseValue receives a string and converts it to a Time
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			t, _ := time.Parse("2006-01-02T15:04:05.000Z", value)
			fmt.Println("Parse time", value, t.Format(time.RFC3339))
			return t
		default:
			fmt.Println("Parse time not string", value)
			return nil
		}
	},
	// ParseLiteral receives an ast.Value and converts it to a Time
	ParseLiteral: func(valueAST ast.Value) interface{} {
		if valueAST, ok := valueAST.(*ast.StringValue); ok {
			t, _ := time.Parse("2006-01-02T15:04:05.000Z", valueAST.Value)
			fmt.Println("Parse time 2", valueAST.Value, t.Format(time.RFC3339))
			return t
		}
		fmt.Println("Parse time not string", valueAST)
		return nil
	},
})
