package lru

import "testing"

func Test_LRU(t *testing.T) {
	l, err := NewLRUWithOption(WithLength(3))
	if err != nil {
		t.Error("new lru err")
	}
	err = l.Put("1", "abc")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru Put err %v", err)
	}
	err = l.Put("2", "abcd")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru Put err %v", err)
	}
	err = l.Put("3", "abcde")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru Put err %v", err)
	}
	v, err := l.Get("2")
	if (err != nil && err != ErrLRUKeyNotFound) || v.(string) != "abcd" {
		t.Errorf("lru Get err %v", err)
	}
	v, err = l.Get("3")
	if (err != nil && err != ErrLRUKeyNotFound) || v.(string) != "abcde" {
		t.Errorf("lru Get err %v", err)
	}
	err = l.Put("4", "ab")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru Put err %v", err)
	}
	v, err = l.Get("1")
	if err != ErrLRUKeyNotFound {
		t.Errorf("lru delete old key err %v, v %v", err, v)
	}
}
