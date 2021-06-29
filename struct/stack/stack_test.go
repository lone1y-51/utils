package stack

import (
	"reflect"
	"testing"
)

func Test_Stack(t *testing.T) {
	q := NewStackWithOptions(
		WithLength(3),
	)
	err := q.Push(1)
	if err != nil {
		t.Error("stack push err")
	}
	err = q.Push(4)
	if err != nil {
		t.Error("stack push err")
	}
	result, err := q.Pop()
	rType := reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 4 {
		t.Error("stack pop err")
	}
	err = q.Push(10)
	if err != nil {
		t.Error("stack push err")
	}
	err = q.Push(11)
	if err != nil {
		t.Error("stack push err")
	}
	err = q.Push(12)
	if err == nil || err != ErrStackIsFull {
		t.Error("stack push err")
	}
	result, err = q.Pop()
	rType = reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 11 {
		t.Error("stack pop err")
	}
	result, err = q.Pop()
	rType = reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 10 {
		t.Error("stack pop err")
	}
	result, err = q.Pop()
	rType = reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 1 {
		t.Error("stack pop err")
	}
	_, err = q.Pop()
	if err == nil || err != ErrStackIsEmpty {
		t.Error("stack pop err")
	}
}
