package http

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

func decode(body io.Reader, v any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(v)
	if err != nil {
		return fmt.Errorf("failed to decode the request body: %w", err)
	}

	validate := validator.New()
	err = validate.Struct(v)
	if err != nil {
		return fmt.Errorf("failed to validate the request body: %w", err)
	}

	return nil
}
