package internal

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrNotAuthorized = errors.New("not authorized")
	ErrDuplicatedKey = errors.New("duplicate entity")
	ErrBadRequest    = errors.New("bad request")
	ErrCacheMiss     = errors.New("cache miss")
	ErrRateLimit     = errors.New("rate limit")
	ErrInternal      = errors.New("server error")
	// etc
)

type WrappedError struct {
	Context string
	Err     error
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

func (e *WrappedError) Unwrap() error {
	return e.Err
}

func WrapErrorMessage(err error, msg string) error {
	return &WrappedError{
		Context: msg,
		Err:     err,
	}
}

func GetErrorContext(err error) string {
	var wrappedErr *WrappedError
	if errors.As(err, &wrappedErr) {
		return wrappedErr.Context
	}
	return err.Error()
}
