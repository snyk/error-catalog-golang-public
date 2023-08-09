package snyk_errors

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalToJSONAPIError(t *testing.T) {
	err := Error{
		ID:         "id",
		Type:       "type",
		Title:      "title",
		StatusCode: 1,
		ErrorCode:  "error-code",
		Level:      "level",
		Detail:     "detail",
		Meta: map[string]any{
			"foo": "bar",
		},
	}

	var buf bytes.Buffer

	require.NoError(t, err.MarshalToJSONAPIError(&buf, "instance"))

	var actual errorsPayload

	require.NoError(t, json.Unmarshal(buf.Bytes(), &actual))

	expected := errorsPayload{
		JSONAPI: errorsPayloadVersion{
			Version: "1.0",
		},
		Errors: []errorObject{
			{
				ID:     "id",
				Title:  "title",
				Detail: "detail",
				Status: "1",
				Code:   "error-code",
				Meta: map[string]any{
					"foo": "bar",
				},
				Links: errorObjectLink{
					About: "type",
				},
				Source: errorObjectSource{
					Pointer: "instance",
				},
			},
		},
	}

	require.Equal(t, expected, actual)
}

type mockWriter func(b []byte) (int, error)

func (w mockWriter) Write(b []byte) (int, error) {
	return w(b)
}

func TestMarshalToJSONAPIErrorWithWriterError(t *testing.T) {
	var err Error

	w := mockWriter(func(b []byte) (int, error) {
		return 0, errors.New("something went wrong")
	})

	require.ErrorContains(t, err.MarshalToJSONAPIError(w, "instance"), "something went wrong")
}
