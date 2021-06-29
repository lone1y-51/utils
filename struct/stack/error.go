package stack

import "errors"

var (
	ErrStackIsFull  = errors.New("stack is full")
	ErrStackIsEmpty = errors.New("stack is empty")
)
