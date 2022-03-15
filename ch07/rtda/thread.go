package rtda

import "jvmgo/ch07/heap"

/**
https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-2.html#jvms-2.5
*/
type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}
func (t *Thread) PC() int {
	return t.pc
}
func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}
func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) TopFrame() *Frame {
	return t.stack.top()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}
func (t *Thread) NewFrame(method *heap.Method) *Frame{
	return NewFrame(t,method)
}

func (t *Thread) IsStackEmpty() bool{
	return t.stack.isEmpty()
}
