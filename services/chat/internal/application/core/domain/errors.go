package domain

import "errors"

var (
	ErrInvalidUserIDs = errors.New("invalid user ids")
	ErrRecordNotFound = errors.New("record not found")
)
