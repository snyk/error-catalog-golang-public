/*
 * © 2026 Snyk Limited
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package snyk_errors

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalToJSONAPIError(t *testing.T) {
	type test struct {
		description string
		error       Error
		expected    jsonAPIDoc
	}

	tests := []test{
		{
			description: "standard error",
			error: Error{
				ID:             "id",
				Type:           "type",
				Title:          "title",
				StatusCode:     1,
				ErrorCode:      "error-code",
				Level:          "level",
				Detail:         "detail",
				Classification: "ACTIONABLE",
				Meta: map[string]any{
					"foo": "bar",
				},
			},
			expected: jsonAPIDoc{
				JSONAPI: jsonAPIObject{
					Version: "1.0",
				},
				Errors: []jsonAPIError{
					{
						ID:     "id",
						Title:  "title",
						Detail: "detail",
						Status: "1",
						Code:   "error-code",
						Meta: map[string]any{
							"foo":                 "bar",
							"level":               "level",
							"classification":      "ACTIONABLE",
							"isErrorCatalogError": true,
						},
						Links: jsonAPILinks{
							About: "type",
						},
						Source: jsonAPIErrSource{
							Pointer: "instance",
						},
					},
				},
			},
		},
		{
			description: "with cause, should not be displayed",
			error: Error{
				ID:             "id",
				Type:           "type",
				Title:          "title",
				StatusCode:     1,
				ErrorCode:      "error-code",
				Level:          "level",
				Detail:         "detail",
				Classification: "ACTIONABLE",
				Cause:          errors.New("something bad"),
				Meta: map[string]any{
					"foo": "bar",
				},
			},
			expected: jsonAPIDoc{
				JSONAPI: jsonAPIObject{
					Version: "1.0",
				},
				Errors: []jsonAPIError{
					{
						ID:     "id",
						Title:  "title",
						Detail: "detail",
						Status: "1",
						Code:   "error-code",
						Meta: map[string]any{
							"foo":                 "bar",
							"level":               "level",
							"classification":      "ACTIONABLE",
							"isErrorCatalogError": true,
						},
						Links: jsonAPILinks{
							About: "type",
						},
						Source: jsonAPIErrSource{
							Pointer: "instance",
						},
					},
				},
			},
		},
		{
			description: "with logs",
			error: Error{
				ID:             "id",
				Type:           "type",
				Title:          "title",
				StatusCode:     1,
				ErrorCode:      "error-code",
				Level:          "warn",
				Detail:         "detail",
				Classification: "ACTIONABLE",
				Logs:           []string{"a", "b"},
				Meta: map[string]any{
					"foo": "bar",
				},
			},
			expected: jsonAPIDoc{
				JSONAPI: jsonAPIObject{
					Version: "1.0",
				},
				Errors: []jsonAPIError{
					{
						ID:     "id",
						Title:  "title",
						Detail: "detail",
						Status: "1",
						Code:   "error-code",
						Meta: map[string]any{
							"foo":                 "bar",
							"level":               "warn",
							"classification":      "ACTIONABLE",
							"isErrorCatalogError": true,
							"logs":                "[\"a\",\"b\"]",
						},
						Links: jsonAPILinks{
							About: "type",
						},
						Source: jsonAPIErrSource{
							Pointer: "instance",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var buf bytes.Buffer
			require.NoError(t, tt.error.MarshalToJSONAPIError(&buf, "instance"))

			var actual jsonAPIDoc
			require.NoError(t, json.Unmarshal(buf.Bytes(), &actual))

			require.Equal(t, tt.expected, actual, "Actual and Expected should be the same for %s", tt.description)
		})
	}
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

func TestMarshalFromJSONAPIError(t *testing.T) {
	expected := Error{
		ID:             "id",
		Type:           "type",
		Title:          "title",
		StatusCode:     1,
		Level:          "level",
		ErrorCode:      "error-code",
		Detail:         "detail",
		Classification: "ACTIONABLE",
		Meta: map[string]any{
			"foo": "bar",
		},
	}

	var buf bytes.Buffer
	require.NoError(t, expected.MarshalToJSONAPIError(&buf, "instance"))

	var actual jsonAPIDoc
	require.NoError(t, json.Unmarshal(buf.Bytes(), &actual))

	first := actual.MarshalFromJSONAPIError()[0]
	require.Equal(t, expected, first)
}

func TestFromJSONAPIErrorBytes(t *testing.T) {
	testDoc := jsonAPIDoc{
		JSONAPI: jsonAPIObject{},
		Errors: []jsonAPIError{
			{
				ID:     "id",
				Title:  "title",
				Status: "1",
				Code:   "SNYK-1234",
				Detail: "detail",
				Meta: map[string]any{
					"foo":                 "bar",
					"level":               "warn",
					"classification":      "ACTIONABLE",
					"isErrorCatalogError": true,
				},
			},
			{
				ID:     "id",
				Title:  "title",
				Status: "1",
				Code:   "SNYK-1234",
				Detail: "detail",
				Meta: map[string]any{
					"foo": "bar",
				},
			},
		},
	}

	expected := []Error{
		{
			ID:             "id",
			Title:          "title",
			StatusCode:     1,
			ErrorCode:      "SNYK-1234",
			Detail:         "detail",
			Level:          "warn",
			Classification: "ACTIONABLE",
			Meta: map[string]any{
				"foo":                 "bar",
				"level":               "warn",
				"classification":      "ACTIONABLE",
				"isErrorCatalogError": true,
			},
		},
		{
			ID:             "id",
			Title:          "title",
			StatusCode:     1,
			ErrorCode:      "SNYK-1234",
			Detail:         "detail",
			Level:          "",
			Classification: "",
			Meta: map[string]any{
				"foo": "bar",
			},
		},
	}

	data, err := json.Marshal(testDoc)
	require.NoError(t, err)

	errors, err := FromJSONAPIErrorBytes(data)
	require.NoError(t, err)
	require.NotNil(t, errors)

	require.Equal(t, expected, errors)
}
