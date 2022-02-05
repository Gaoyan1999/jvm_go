package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
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
