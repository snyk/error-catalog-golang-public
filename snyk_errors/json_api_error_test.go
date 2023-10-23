/*
 * © 2023 Snyk Limited
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
	err := Error{
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
				ID:             "id",
				Title:          "title",
				Detail:         "detail",
				Status:         "1",
				Code:           "error-code",
				Classification: "ACTIONABLE",
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
