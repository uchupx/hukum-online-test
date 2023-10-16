package errors

import "errors"

var (
	ErrDataIsExist  = errors.New("data is exist")
	ErrDataNotFound = errors.New("data not found")
)
