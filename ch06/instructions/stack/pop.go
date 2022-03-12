package stack

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

func (ins *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopSlot()
}
func (ins *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopSlot()
	frame.OperandStack.PopSlot()
}
