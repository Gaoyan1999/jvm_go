package rtda

import "jvmgo/ch08/heap"

/**
https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-2.html#jvms-2.5
*/
type Thread struct {
	pc    int
	Stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		Stack: newStack(1024),
	}
}
func (t *Thread) PC() int {
	return t.pc
}
func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.Stack.push(frame)
}
func (t *Thread) PopFrame() *Frame {
	return t.Stack.pop()
}

func (t *Thread) TopFrame() *Frame {
	return t.Stack.top()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.Stack.top()
}
func (t *Thread) NewFrame(method *heap.Method) *Frame{
	return NewFrame(t,method)
}

func (t *Thread) IsStackEmpty() bool{
	return t.Stack.isEmpty()
}
