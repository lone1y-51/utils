package dlink

import "errors"

var (
	ErrDLinkIsFull  = errors.New("dlink is full")
	ErrDLinkIsEmpty = errors.New("dlink is empty")
)
