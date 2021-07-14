package lru

import "testing"

func Test_LRU(t *testing.T) {
	l, err := NewLRUWithOption(WithLength(3))
	if err != nil {
		t.Error("new lru err")
	}
	err = l.PUT("1", "abc")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru put err %v", err)
	}
	err = l.PUT("2", "abcd")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru put err %v", err)
	}
	err = l.PUT("3", "abcde")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru put err %v", err)
	}
	v, err := l.GET("2")
	if (err != nil && err != ErrLRUKeyNotFound) || v.(string) != "abcd" {
		t.Errorf("lru get err %v", err)
	}
	v, err = l.GET("3")
	if (err != nil && err != ErrLRUKeyNotFound) || v.(string) != "abcde" {
		t.Errorf("lru get err %v", err)
	}
	err = l.PUT("4", "ab")
	if err != nil && err != ErrLRUIsFull {
		t.Errorf("lru put err %v", err)
	}
	v, err = l.GET("1")
	if err != ErrLRUKeyNotFound {
		t.Errorf("lru delete old key err %v, v %v", err, v)
	}
}
