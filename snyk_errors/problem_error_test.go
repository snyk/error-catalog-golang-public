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
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorMeta(t *testing.T) {
	var err Error

	WithMeta("foo", "bar")(&err)

	if err.Meta == nil {
		t.Fatalf("meta not initiliazed")
	}

	if v := err.Meta["foo"]; v != "bar" {
		t.Fatalf("invalid meta value: %v", v)
	}
}

func TestErrorCause(t *testing.T) {
	var cause = errors.New("cause")

	var err Error

	WithCause(cause)(&err)

	if !errors.Is(err, cause) {
		t.Fatalf("not wrapping the cause")
	}
}

func TestErrorLogs(t *testing.T) {
	var logs = []string{"a", "b"}

	var err Error

	WithLogs(logs)(&err)

	assert.ElementsMatch(t, logs, err.Logs)
}
