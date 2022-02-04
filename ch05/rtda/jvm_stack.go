package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if stack._top != nil {
		frame.next = stack._top
	}
	stack._top = frame
	stack.size++
}
func (stack *Stack) pop() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}
	pop := stack._top
	stack._top.next = nil
	stack._top = stack._top.next
	stack.size--
	return pop
}
func (stack *Stack) top() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}
	return stack._top
}
