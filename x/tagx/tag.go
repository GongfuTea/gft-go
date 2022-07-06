package tagx

import (
	"reflect"
	"strings"
)

type Tag struct {
	Key  string
	Name string
}

type Tags struct {
	field reflect.StructField
}

func Parse(field reflect.StructField) *Tags {
	return &Tags{
		field: field,
	}
}

func (t *Tags) GetJsonName() string {
	fieldName := t.field.Name
	switch jsonTag := t.field.Tag.Get("json"); jsonTag {
	case "-":
	case "":
		return fieldName
	default:
		parts := strings.Split(jsonTag, ",")
		name := parts[0]
		if name == "" {
			return fieldName
		}
		return name
	}
	return fieldName
}

func GetJsonName(field reflect.StructField) string {
	return Parse(field).GetJsonName()
}
