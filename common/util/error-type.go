package util

import (
	"errors"
	"net/http"
)

var ErrHasing = errors.New("error in hashing the password")
var ErrInternal = errors.New("internal error")
var CustomErrorType = map[error]int{
	ErrHasing:   http.StatusInternalServerError,
	ErrInternal: http.StatusInternalServerError,
}
