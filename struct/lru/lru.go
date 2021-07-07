package lru

import (
	"github.com/lone1y-51/utils/struct/dlink"
	"github.com/pkg/errors"
)

type LRU struct {
	DLink  *dlink.DLink
	Length int
	Cache  map[interface{}]*dlink.DLinkNode
}

func NewLRUCache(length int) (*LRU, error) {
	link, err := dlink.NewDLink(length)
	if err != nil {
		return nil, errors.Wrap(err, "new lru err")
	}
	cacheMap := make(map[interface{}]*dlink.DLinkNode, length)
	l := &LRU{
		DLink:  link,
		Length: length,
		Cache:  cacheMap,
	}
	return l, nil
}

func (l *LRU) PUT(key, value interface{}) error {
	if node, ok := l.Cache[key]; ok {
		node.Value = value
		_ = l.DLink.MoveNodeToHead(node)
		return nil
	}
	if l.DLink.IsFull() {
		_ = l.DLink.RemoveTail()
	}
	node, err := l.DLink.AddToHead(value)
	if err != nil {
		return err
	}
	l.Cache[key] = node
	return nil
}

func (l *LRU) GET(key interface{}) (interface{}, error) {
	if node, ok := l.Cache[key]; ok {
		_ = l.DLink.MoveNodeToHead(node)
		return node.Value, nil
	}
	return nil, errors.New("key not found")
}
