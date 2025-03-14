package stockxgo

import (
	"errors"
)

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrBadRequest    = errors.New("bad request")
	ErrInternal      = errors.New("internal server error")
	ErrUnknownStatus = errors.New("unknown status code")
)

func statusCode(statusCode int) error {
	switch statusCode {
	case 200:
		return nil
	case 401:
		return ErrUnauthorized
	case 400:
		return ErrBadRequest
	case 500:
		return ErrInternal
	default:
		return ErrUnknownStatus
	}
}
