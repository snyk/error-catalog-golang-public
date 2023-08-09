package snyk_errors

import (
	"errors"
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
