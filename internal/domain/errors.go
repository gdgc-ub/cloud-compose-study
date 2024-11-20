package domain

import "errors"

var (
	ErrNotFound = errors.New("item not found")
)

func GetCode(err error) int {
	if err == nil {
		return 200
	}

	switch err {
	case ErrNotFound:
		return 404
	default:
		return 500 
	}
}