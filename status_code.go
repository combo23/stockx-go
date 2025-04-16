package stockxgo

import (
	"errors"
	"fmt"
)

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrBadRequest    = errors.New("bad request")
	ErrInternal      = errors.New("internal server error")
	ErrUnknownStatus = errors.New("unknown status code: %v")
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
		return fmt.Errorf(ErrUnknownStatus.Error(), statusCode)
	}
}
