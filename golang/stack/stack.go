package stack

const (
	overflow  = 1001
	underflow = 0110
	undefined = 0
)

type Stack struct {
	maxSize  int
	size     int
	elements []int
}

func MakeStack(size int) Stack {
	return Stack{
		maxSize:  size,
		size:     0,
		elements: make([]int, size),
	}
}

func (s *Stack) Pop() (val int, err int) {

	err = undefined
	if s.size <= 0 {
		err = underflow
	} else {
		s.size--
		val = s.elements[s.size]
	}

	return val, err
}

func (s *Stack) Push(val int) (err int) {
	err = undefined

	if s.size >= s.maxSize {
		err = overflow
	} else {
		s.elements[s.size] = val
		s.size++
	}

	return err
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) MaxSize() int {
	return s.maxSize
}
