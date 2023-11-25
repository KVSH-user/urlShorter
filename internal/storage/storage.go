package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url nit found")
	ErrURLExists   = errors.New("url exists")
)
