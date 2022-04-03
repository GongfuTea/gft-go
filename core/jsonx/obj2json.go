package jsonx

import (
	"encoding/json"
	"fmt"
)

func PrintAsJson(data any) {
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-----------------")
	fmt.Println(string(e))
}
