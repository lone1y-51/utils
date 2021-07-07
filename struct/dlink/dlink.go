package dlink

type DLinkNode struct {
	Prev  *DLinkNode
	Next  *DLinkNode
	Value interface{}
}

// DLink is a doubly link
type DLink struct {
	Head   *DLinkNode
	Tail   *DLinkNode
	Length int
	Count  int
}

func NewDLink(length int) (*DLink, error) {
	link := &DLink{
		Length: length,
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
func (l *DLink) Add(value interface{}) (*DLinkNode, error) {
	if l.IsFull() {
		return nil, ErrDLinkIsFull
	}
	node := &DLinkNode{
		Value: value,
	}
	node.Next = l.Tail
	node.Prev = l.Tail.Prev
	l.Tail.Prev.Next = node
	l.Tail.Prev = node
	l.Count++
	return node, nil
}

func (l *DLink) AddToHead(value interface{}) (*DLinkNode, error) {
	if l.IsFull() {
		return nil, ErrDLinkIsFull
	}
	node := &DLinkNode{
		Value: value,
	}
	node.Prev = l.Head
	node.Next = l.Head.Next
	l.Head.Next.Prev = node
	l.Head.Next = node
	l.Count++
	return node, nil
}

func (l *DLink) RemoveTail() error {
	if l.IsEmpty() {
		return ErrDLinkIsEmpty
	}
	tailNode := l.Tail.Prev
	l.removeNode(tailNode)
	return nil
}

func (l *DLink) MoveNodeToHead(node *DLinkNode) error {
	l.removeNode(node)
	l.moveToHead(node)
	return nil
}

func (l *DLink) removeNode(node *DLinkNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (l *DLink) moveToHead(node *DLinkNode) {
	node.Prev = l.Head
	node.Next = l.Head.Next
	l.Head.Next.Prev = node
	l.Head.Next = node
}
