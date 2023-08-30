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
	"encoding/json"
	"io"
	"strconv"

	_ "github.com/google/uuid"
)

type errorsPayload struct {
	JSONAPI errorsPayloadVersion `json:"jsonapi"`
	Errors  []errorObject        `json:"errors"`
}

type errorsPayloadVersion struct {
	Version string `json:"version"`
}

type errorObject struct {
	ID     string                 `json:"id,omitempty"`
	Title  string                 `json:"title,omitempty"`
	Detail string                 `json:"detail,omitempty"`
	Status string                 `json:"status,omitempty"`
	Code   string                 `json:"code,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Links  errorObjectLink        `json:"links,omitempty"`
	Source errorObjectSource      `json:"source"`
}

type errorObjectLink struct {
	About string `json:"about,omitempty"`
}

type errorObjectSource struct {
	Pointer string `json:"pointer,omitempty"`
}

func (e Error) MarshalToJSONAPIError(w io.Writer, instance string) error {
	payload := errorsPayload{
		JSONAPI: errorsPayloadVersion{
			Version: "1.0",
		},
		Errors: []errorObject{
			{
				ID:     e.ID,
				Title:  e.Title,
				Detail: e.Detail,
				Status: strconv.Itoa(e.StatusCode),
				Code:   e.ErrorCode,
				Meta:   e.Meta,
				Links: errorObjectLink{
					About: e.Type,
				},
				Source: errorObjectSource{
					Pointer: instance,
				},
			},
		},
	}
	return json.NewEncoder(w).Encode(payload)
}
