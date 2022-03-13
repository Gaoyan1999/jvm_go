package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type InvokeSpecial struct{ base.Index16Instruction }
// hack!
func (ins *InvokeSpecial) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopRef()
}