package jsonx

import (
	"encoding/json"
	"fmt"
)

func PrintAsJson(label string, data any) {
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(label, err)
		return
	}
	fmt.Println(label, string(e))
}
