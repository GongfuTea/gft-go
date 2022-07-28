package jsonx

import (
	"encoding/json"
	"io"
)

func Encode(v any) io.Reader {
	reader, writer := io.Pipe()
	go func() {
		err := json.NewEncoder(writer).Encode(v)
		if err != nil {
			writer.CloseWithError(err)
		} else {
			writer.Close()
		}
	}()
	return reader
}
