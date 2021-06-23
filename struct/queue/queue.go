package queue

type Option struct {
	Length int
}

type Queue struct {
	data   []interface{}
	length int
}

type OptFunc func(*Option)

func newQueue(option *Option) *Queue {
	return &Queue{
		length: option.Length,
	}
}

func WithQueueLength(length int) OptFunc {
	return func(opt *Option) {
		opt.Length = length
	}
}

func NewQueueWithOptions(opts ...OptFunc) *Queue {
	option := &Option{}
	for _, opt := range opts {
		opt(option)
	}
	return newQueue(option)
}

func (q *Queue) Push(d interface{}) error {
	if q.length != 0 && len(q.data) == q.length {
		return ErrQueueIsFull
	}
	q.data = append(q.data, d)
	return nil
}

func (q *Queue) Pop() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, ErrQueueIsEmpty
	}
	result := q.data[0]
	q.data = q.data[1:len(q.data)]
	return result, nil
}

func (q *Queue) Count() int {
	return len(q.data)
}
