package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type InvokeSpecial struct{ base.Index16Instruction }
// hack!
func (ins *InvokeSpecial) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopRef()
}