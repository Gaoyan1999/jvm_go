package rtda

import "jvmgo/ch06/heap"

/**
https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-2.html#jvms-2.6

A frame is used to store data and partial results, as well as to perform dynamic linking, return values for methods, and dispatch exceptions.
A new frame is created each time a method is invoked.
A frame is destroyed when its method invocation completes,
whether that completion is normal or abrupt (it throws an uncaught exception).
Frames are allocated from the Java Virtual Machine stack (§2.5.2) of the thread creating the frame.
Each frame has its own array of local variables (§2.6.1), its own operand stack (§2.6.2),
and a reference to the run-time constant pool (§2.5.5) of the class of the current method.

*/
type Frame struct {
	LocalVars
	*OperandStack
	next      *Frame
	Thread    *Thread
	Method    *heap.Method
	maxLocals uint
	maxStack  uint
	NextPC    int // the next instruction after the call
}

func NewFrame(t *Thread, method *heap.Method, maxLocal uint, maxStack uint) *Frame {
	return &Frame{
		Thread:       t,
		Method:       method,
		LocalVars:    newLocalVars(maxLocal),
		OperandStack: newOperandStack(maxStack),
	}
}
