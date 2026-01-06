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

type Error struct {
	ID             string         `json:"id,omitempty"`
	Type           string         `json:"type,omitempty"`
	Title          string         `json:"title,omitempty"`
	StatusCode     int            `json:"statusCode,omitempty"`
	ErrorCode      string         `json:"errorCode,omitempty"`
	Description    string         `json:"description,omitempty"`
	Level          string         `json:"level,omitempty"`
	Links          []string       `json:"links,omitempty"`
	Detail         string         `json:"detail,omitempty"`
	Meta           map[string]any `json:"meta,omitempty"`
	Cause          error          `json:"cause,omitempty"`
	Classification string         `json:"classification,omitempty"`
	Logs           []string       `json:"logs,omitempty"`
}

func (e Error) Error() string {
	return e.Title
}

func (e Error) Unwrap() error {
	return e.Cause
}

var _ error = Error{}

type Option func(e *Error)

func WithMeta(key string, value any) Option {
	return func(e *Error) {
		if e.Meta == nil {
			e.Meta = make(map[string]any)
		}

		e.Meta[key] = value
	}
}

func WithCause(cause error) Option {
	return func(e *Error) {
		e.Cause = cause
	}
}

func WithLogs(logs []string) Option {
	return func(e *Error) {
		e.Logs = logs
	}
}

func WithLinks(links []string) Option {
	return func(e *Error) {
		e.Links = append(e.Links, links...)
	}
}
