package dlink

type Options struct {
	length int
}

type DLinkNode struct {
	Prev  *DLinkNode
	Next  *DLinkNode
	Value any
	Key   any
}

// DLink is a doubly link
type DLink struct {
	Head   *DLinkNode
	Tail   *DLinkNode
	Length int
	Count  int
}

type OptFunc func(opt *Options)

func NewDLinkWithOptions(optFuncs ...OptFunc) (*DLink, error) {
	opt := &Options{}
	for _, optFunc := range optFuncs {
		optFunc(opt)
	}
	return newDLink(opt)
}

func WithLength(length int) OptFunc {
	return func(opt *Options) {
		opt.length = length
	}
}

func newDLink(opt *Options) (*DLink, error) {
	link := &DLink{
		Length: opt.length,
		Head:   &DLinkNode{},
		Tail:   &DLinkNode{},
		Count:  0,
	}
	link.Head.Next = link.Tail
	link.Tail.Prev = link.Head
	return link, nil
}

func (l *DLink) IsFull() bool {
	return l.Count >= l.Length
}

func (l *DLink) IsEmpty() bool {
	return l.Count <= 0
}

// Add to link tail
func (l *DLink) Add(key, value any) (*DLinkNode, error) {
	if l.IsFull() {
		return nil, ErrDLinkIsFull
	}
	node := &DLinkNode{
		Value: value,
		Key:   key,
	}
	node.Next = l.Tail
	node.Prev = l.Tail.Prev
	l.Tail.Prev.Next = node
	l.Tail.Prev = node
	l.Count++
	return node, nil
}

func (l *DLink) AddToHead(key, value any) (*DLinkNode, error) {
	if l.IsFull() {
		return nil, ErrDLinkIsFull
	}
	node := &DLinkNode{
		Value: value,
		Key:   key,
	}
	node.Prev = l.Head
	node.Next = l.Head.Next
	l.Head.Next.Prev = node
	l.Head.Next = node
	l.Count++
	return node, nil
}

func (l *DLink) RemoveTail() (*DLinkNode, error) {
	if l.IsEmpty() {
		return nil, ErrDLinkIsEmpty
	}
	tailNode := l.Tail.Prev
	l.removeNode(tailNode)
	l.Count--
	return tailNode, nil
}

func (l *DLink) MoveNodeToHead(node *DLinkNode) error {
	l.removeNode(node)
	l.moveToHead(node)
	return nil
}

func (l *DLink) removeNode(node *DLinkNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	node.Prev = nil
	node.Next = nil
}

func (l *DLink) moveToHead(node *DLinkNode) {
	node.Prev = l.Head
	node.Next = l.Head.Next
	l.Head.Next.Prev = node
	l.Head.Next = node
}
