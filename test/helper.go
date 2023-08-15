package test

import (
	"encoding/json"
	"io"
)

func ReadResponseBody(respBody io.Reader, responseObject any) error {
	body, err := io.ReadAll(respBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		return err
	}

	return nil
}
