package stack

type Option struct {
	length int
}

type OptFunc func(opt *Option)

type Stack struct {
	data   []interface{}
	length int
}

func newStack(opt *Option) *Stack {
	s := &Stack{}
	s.length = opt.length
	s.data = make([]interface{}, 0)
	return s
}

func WithLength(length int) OptFunc {
	return func(opt *Option) {
		opt.length = length
	}
}

func NewStackWithOptions(options ...OptFunc) *Stack {
	opt := &Option{}
	for _, optFunc := range options {
		optFunc(opt)
	}
	return newStack(opt)
}

func (s *Stack) Count() int {
	return len(s.data)
}

func (s *Stack) Push(data interface{}) error {
	if len(s.data) == s.length && s.length != 0 {
		return ErrStackIsFull
	}
	s.data = append(s.data, data)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, ErrStackIsEmpty
	}
	index := len(s.data) - 1
	result := s.data[index]
	s.data = s.data[:index]
	return result, nil
}

func (s *Stack) Clear() {
	s.data = make([]interface{}, 0)
}
