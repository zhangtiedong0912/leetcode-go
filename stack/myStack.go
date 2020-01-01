package stack

type MyStack struct {
	data     []interface{}
	length   int
}

func InitMyStack() *MyStack {
	return &MyStack{data: []interface{}{}, length: 0}
}

func (s *MyStack) Push(data interface{}) {
	s.data = append(s.data,data)
	s.length++
}

func (s *MyStack) Pop() interface{} {
	if s.length <= 0 {
		panic("int stack pop: index out of range")
	}

	t := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return t
}

func (s *MyStack) Peek() interface{} {
	if s.length <= 0 {
		panic("empty stack")
	}

	return s.data[s.length-1]
}

func (s *MyStack) IsEmpty() bool {
	b := s.length == 0
	return b
}
