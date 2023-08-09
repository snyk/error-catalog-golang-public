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
