package snyk_errors

type Error struct {
	ID         string
	Type       string
	Title      string
	StatusCode int
	ErrorCode  string
	Level      string
	Links      []string
	Detail     string
	Meta       map[string]any
	Cause      error
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
