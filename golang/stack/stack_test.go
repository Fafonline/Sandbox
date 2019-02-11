package stack

import (
	"testing"
)

func TestWhenCreateStackWithMaxSizeX_MaxSizeIsX(t *testing.T) {

	maxSize := 10
	stack := MakeStack(maxSize)

	size := stack.MaxSize()
	wants := maxSize

	if size != wants {
		t.Errorf("Expect %v. Got %v", wants, size)
	}
}

func TestWhenPopInEmpty_ReturnUnderflowError(t *testing.T) {
	stack := MakeStack(0)

	_, err := stack.Pop()

	wants := underflow

	if err != wants {
		t.Errorf("Expect %v. Got %v", wants, err)
	}
}

func TestWhenPushInFull_ReturnOverflowError(t *testing.T) {
	stack := MakeStack(0)

	err := stack.Push(0)

	wants := overflow

	if err != wants {
		t.Errorf("Expect %v. Got %v", wants, err)
	}
}

func TestWhenPushInEmpty_SizeIsOne(t *testing.T) {
	stack := MakeStack(1)

	stack.Push(0)
	size := stack.Size()

	wants := 1

	if size != wants {
		t.Errorf("Expect %v. Got %v", wants, size)
	}
}

func GivenANonEmptyStackWith5Element() Stack {
	stack := MakeStack(10)
	i := 0
	for i < 5 {
		stack.Push(0)
		i++
	}

	return stack
}

func TestWhenPopNonEmptyStackWith5Element_SizeIs4(t *testing.T) {
	stack := GivenANonEmptyStackWith5Element()

	stack.Pop()
	size := stack.Size()

	wants := 4

	if size != wants {
		t.Errorf("Expect %v. Got %v", wants, size)
	}
}

func TestWhenPoppedTooMuchFromNonEmpty_Returnunderflow(t *testing.T) {

	type testCase struct {
		name  string
		given int
		want  int
	}

	tests := []testCase{
		{
			name: "Pop #1",
			want: undefined,
		},
		{
			name: "Pop #2",
			want: undefined,
		},
		{
			name: "Pop #3",
			want: undefined,
		},
		{
			name: "Pop #4",
			want: undefined,
		},
		{
			name: "Pop #5",
			want: undefined,
		},
		{
			name: "Pop #6",
			want: underflow,
		},
	}

	stack := GivenANonEmptyStackWith5Element()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := stack.Pop()
			if got != tt.want {
				t.Errorf("%s got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
func TestWhenPushTooMuchFromNonEmpty_ReturnOverflow(t *testing.T) {
	type testCase struct {
		name  string
		given int
		want  int
	}

	tests := []testCase{
		{
			name: "Push #1",
			want: undefined,
		},
		{
			name: "Push #2",
			want: undefined,
		},
		{
			name: "Push #3",
			want: undefined,
		},
		{
			name: "Push #4",
			want: undefined,
		},
		{
			name: "Push #5",
			want: undefined,
		},
		{
			name: "Push #6",
			want: overflow,
		},
	}

	stack := GivenANonEmptyStackWith5Element()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stack.Push(tt.given)
			if got != tt.want {
				t.Errorf("%s got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestWhenPushNonEmptyStackWith5Element_SizeIs4(t *testing.T) {
	stack := GivenANonEmptyStackWith5Element()

	stack.Push(0)
	size := stack.Size()

	wants := 6

	if size != wants {
		t.Errorf("Expect %v. Got %v", wants, size)
	}
}

func TestWhenPushInNonEmptyStack_PoppedValueIsLastlyPushed(t *testing.T) {
	stack := GivenANonEmptyStackWith5Element()

	stack.Push(0xDEAD)
	val, _ := stack.Pop()

	wants := 0xDEAD

	if val != wants {
		t.Errorf("Expect %v. Got %v", wants, val)
	}
}

func TestWhenPushSeveralTimeInEmptyStack_PoppedValueIsLastlyPushed(t *testing.T) {
	stack := MakeStack(10)
	stack.Push(0xDEAD)
	stack.Push(0xCAFE)

	val, _ := stack.Pop()
	wants := 0xCAFE
	if val != wants {
		t.Errorf("Expect %v. Got %v", wants, val)
	}
}

func TestWhenPushTimeInEmptyStack_SecondPoppedValueIsTheFirstPushed(t *testing.T) {
	stack := MakeStack(10)
	stack.Push(0xDEAD)
	stack.Push(0xCAFE)
	stack.Pop()
	val, _ := stack.Pop()

	wants := 0xDEAD
	if val != wants {
		t.Errorf("Expect %v. Got %v", wants, val)
	}
}

func TestWhenPushValuesInAnOrder_ValuesArePoppedInReverseOrder(t *testing.T) {
	stack := MakeStack(10)

	stack.Push(0xDEAD)
	stack.Push(0xCAFE)
	stack.Push(0xFFFF)
	val, _ := stack.Pop()
	wants := 0xFFFF
	if val != wants {
		t.Errorf("Expect %v. Got %v", wants, val)
	}
	val, _ = stack.Pop()
	wants = 0xCAFE
	if val != wants {
		t.Errorf("Expect %v. Got %v", wants, val)
	}
	val, _ = stack.Pop()
	wants = 0xDEAD
	if val != wants {
		t.Errorf("Expect %v. Got %v", wants, val)
	}
}
