/*
 * © 2024 Snyk Limited
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
	ID             string
	Type           string
	Title          string
	StatusCode     int
	ErrorCode      string
	Description    string
	Level          string
	Links          []string
	Detail         string
	Meta           map[string]any
	Cause          error
	Classification string
	Logs           []string
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
