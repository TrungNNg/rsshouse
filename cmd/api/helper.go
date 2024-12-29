package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// readJSON reads and decodes a JSON request body into the specified destination `dst`.
// It includes validation to ensure the request body adheres to expected size and format constraints.
//
// Parameters:
// - w: The http.ResponseWriter for sending error responses if necessary.
// - r: The *http.Request containing the JSON body to decode.
// - dst: A pointer to the destination object where the JSON will be unmarshaled.
//
// Returns:
// - An error describing why the JSON decoding failed.
//
// Validates scenarios:
//   - Malformed JSON
//   - Unexpected data types in the JSON
//   - Unknown JSON fields
//   - Empty request bodies
//   - Oversized bodies
//   - Multiple JSON values
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Limit the size of the request body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Initialize json.Decoder.
	dec := json.NewDecoder(r.Body)

	// Decode() will now return error if JSON has unknown fields.
	dec.DisallowUnknownFields()

	// Decode the request body to the destination.
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	// Call Decode() again using struct{}{} which will consume extra JSON data then discard it.
	// If the request body only contained a single JSON value this will
	// return an io.EOF error. If we get any other error or nil then there is extra data
	// so we return our own custom error message.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
