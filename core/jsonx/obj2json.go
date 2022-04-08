package jsonx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func PrintAsJson(label string, data any) {
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(label, err)
		return
	}
	fmt.Println(label, string(e))
}

func ReadJsonByte(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}
