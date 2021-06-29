package queue

import (
	"reflect"
	"testing"
)

func Test_Queue(t *testing.T) {
	q := NewQueueWithOptions(
		WithQueueLength(3),
	)
	err := q.Push(1)
	if err != nil {
		t.Error("queue push err")
	}
	err = q.Push(4)
	if err != nil {
		t.Error("queue push err")
	}
	result, err := q.Pop()
	rType := reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 1 {
		t.Error("queue pop err")
	}
	err = q.Push(10)
	if err != nil {
		t.Error("queue push err")
	}
	err = q.Push(11)
	if err != nil {
		t.Error("queue push err")
	}
	err = q.Push(12)
	if err == nil || err != ErrQueueIsFull {
		t.Error("queue push err")
	}
	result, err = q.Pop()
	rType = reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 4 {
		t.Error("queue pop err")
	}
	result, err = q.Pop()
	rType = reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 10 {
		t.Error("queue pop err")
	}
	result, err = q.Pop()
	rType = reflect.TypeOf(result)
	if err != nil || rType.Kind() != reflect.Int || result.(int) != 11 {
		t.Error("queue pop err")
	}
	_, err = q.Pop()
	if err == nil || err != ErrQueueIsEmpty {
		t.Error("queue pop err")
	}
}
