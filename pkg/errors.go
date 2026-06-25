package pkg

import "errors"

type Errors string

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternal     = errors.New("internal server error")
	ErrForbidden    = errors.New("forbidden")
	ErrNotFound     = errors.New("resource not found")
	ErrBadRequest   = errors.New("bad request")
	ErrConflict     = errors.New("resource already exists")
)
