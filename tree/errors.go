package tree

import "errors"

var (
	ErrEmptyTree = errors.New("tree is empty")
	ErrNotFound  = errors.New("not found dst key")
)
