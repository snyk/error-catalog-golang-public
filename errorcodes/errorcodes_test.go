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
package errorcodes_test

import (
  "github.com/snyk/error-catalog-golang-public/errorcodes"
  "testing"
)

func TestErrorCodeMap(t *testing.T) {
  got := errorcodes.Snyk.TooManyRequestsError
  want := "SNYK-0001"

  if got != want {
    t.Errorf("got %s, wanted %s", got, want)
  }
}
