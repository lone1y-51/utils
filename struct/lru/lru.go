package lru

import (
	"github.com/lone1y-51/utils/struct/dlink"
	"github.com/pkg/errors"
)

type Options struct {
	length int
}

type LRU struct {
	DLink  *dlink.DLink
	Length int
	Cache  map[any]*dlink.DLinkNode
}

type OptFunc func(opt *Options)

func newLRUCache(opt *Options) (*LRU, error) {
	link, err := dlink.NewDLinkWithOptions(dlink.WithLength(opt.length))
	if err != nil {
		return nil, errors.Wrap(err, "new lru err")
	}
	cacheMap := make(map[any]*dlink.DLinkNode, opt.length)
	l := &LRU{
		DLink:  link,
		Length: opt.length,
		Cache:  cacheMap,
	}
	return l, nil
}

func NewLRUWithOption(optFuncs ...OptFunc) (*LRU, error) {
	opt := &Options{}
	for _, optFunc := range optFuncs {
		optFunc(opt)
	}
	return newLRUCache(opt)
}

func WithLength(length int) OptFunc {
	return func(opt *Options) {
		opt.length = length
	}
}

func (l *LRU) Put(key, value any) error {
	if node, ok := l.Cache[key]; ok {
		node.Value = value
		_ = l.DLink.MoveNodeToHead(node)
		return nil
	}
	if l.DLink.IsFull() {
		tail, _ := l.DLink.RemoveTail()
		delete(l.Cache, tail.Key)
	}
	node, err := l.DLink.AddToHead(key, value)
	if err != nil {
		return err
	}
	l.Cache[key] = node
	return nil
}

func (l *LRU) Get(key any) (any, error) {
	if node, ok := l.Cache[key]; ok {
		_ = l.DLink.MoveNodeToHead(node)
		return node.Value, nil
	}
	return nil, ErrLRUKeyNotFound
}
