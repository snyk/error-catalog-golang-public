/*
 * © 2025 Snyk Limited
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

type jsonAPIDoc struct {
	JSONAPI jsonAPIObject  `json:"jsonapi"`
	Errors  []jsonAPIError `json:"errors"`
}

type jsonAPIObject struct {
	Version string `json:"version"`
}

type jsonAPIError struct {
	ID     string                 `json:"id,omitempty"`
	Title  string                 `json:"title,omitempty"`
	Detail string                 `json:"detail,omitempty"`
	Status string                 `json:"status,omitempty"`
	Code   string                 `json:"code,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Links  jsonAPILinks           `json:"links,omitempty"`
	Source jsonAPIErrSource       `json:"source,omitempty"`
}

type jsonAPILinks struct {
	About string `json:"about,omitempty"`
}

type jsonAPIErrSource struct {
	Pointer string `json:"pointer,omitempty"`
}

func (e Error) MarshalToJSONAPIError(w io.Writer, instance string) error {
	err := jsonAPIError{
		ID:     e.ID,
		Title:  e.Title,
		Detail: e.Detail,
		Status: strconv.Itoa(e.StatusCode),
		Code:   e.ErrorCode,
		Meta:   make(map[string]any),
		Links: jsonAPILinks{
			About: e.Type,
		},
		Source: jsonAPIErrSource{
			Pointer: instance,
		},
	}

	if e.Meta != nil {
		err.Meta = e.Meta
	}

	if e.Logs != nil {
		j, _ := json.Marshal(e.Logs)
		err.Meta["logs"] = string(j)
	}

	err.Meta["classification"] = e.Classification

	// Allow consumers to probe if this specific type of JsonApi response originates from an error catalog error.
	err.Meta["isErrorCatalogError"] = true

	return json.NewEncoder(w).Encode(jsonAPIDoc{
		JSONAPI: jsonAPIObject{
			Version: "1.0",
		},
		Errors: []jsonAPIError{err},
	})
}

func (j jsonAPIDoc) MarshalFromJSONAPIError() []Error {
	var errors []Error

	for _, jsonAPIErr := range j.Errors {
		status, _ := strconv.Atoi(jsonAPIErr.Status)

		err := Error{
			ID:             jsonAPIErr.ID,
			Type:           jsonAPIErr.Links.About,
			Title:          jsonAPIErr.Title,
			StatusCode:     status,
			ErrorCode:      jsonAPIErr.Code,
			Detail:         jsonAPIErr.Detail,
			Meta:           jsonAPIErr.Meta,
			Classification: jsonAPIErr.Meta["classification"].(string),
		}

		if err.Logs != nil {
			err.Meta["logs"] = err.Logs
		}

		errors = append(errors, err)
	}

	return errors
}
