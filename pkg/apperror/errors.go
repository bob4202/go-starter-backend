package apperror

import "errors"

var (
	ErrBadRequest     = errors.New("invalid request body")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrInternalServer = errors.New("internal server error")
)
