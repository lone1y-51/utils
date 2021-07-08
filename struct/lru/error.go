package lru

import "github.com/pkg/errors"

var (
	ErrLRUIsFull      = errors.New("lru is full")
	ErrLRUKeyNotFound = errors.New("key not found")
)
