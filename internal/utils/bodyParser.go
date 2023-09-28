package utils

import (
	"encoding/json"
	"io"
)

func BodyParser[T comparable](body io.ReadCloser, parsed *T) error {
	decoder := json.NewDecoder(body)

	err := decoder.Decode(&parsed)
	if err != nil {
		return err
	}

	return nil
}
