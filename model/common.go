package model

import "errors"

var (
	// ErrNotExist represent error when data not exist in database
	ErrNotExist = errors.New("not exist")
)
