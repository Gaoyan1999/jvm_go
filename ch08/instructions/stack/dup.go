package stack

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)
type DUP struct{ base.NoOperandsInstruction }
func (ins *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ base.NoOperandsInstruction }
func (instr *DUP_X1) Execute(frame *rtda.Frame) {
	val1 := frame.PopSlot()
	val2 := frame.PopSlot()
	frame.PushSlot(val1)
	frame.PushSlot(val2)
	frame.PushSlot(val1)
}

// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperandsInstruction }
func (instr *DUP_X2) Execute(frame *rtda.Frame) {
	val1 := frame.PopSlot()
	val2 := frame.PopSlot()
	val3 := frame.PopSlot()
	frame.PushSlot(val1)
	frame.PushSlot(val3)
	frame.PushSlot(val2)
	frame.PushSlot(val1)
}

type DUP2 struct{ base.NoOperandsInstruction }
func (instr *DUP2) Execute(frame *rtda.Frame) {
	val1 := frame.PopSlot()
	val2 := frame.PopSlot()
	frame.PushSlot(val2)
	frame.PushSlot(val1)
	frame.PushSlot(val2)
	frame.PushSlot(val1)
}
type DUP2_X1 struct{ base.NoOperandsInstruction }

func (instr *DUP2_X1) Execute(frame *rtda.Frame) {
	val1 := frame.PopSlot()
	val2 := frame.PopSlot()
	val3 := frame.PopSlot()
	frame.PushSlot(val2)
	frame.PushSlot(val1)
	frame.PushSlot(val3)
	frame.PushSlot(val2)
	frame.PushSlot(val1)
}
type DUP2_X2 struct{ base.NoOperandsInstruction }
func (instr *DUP2_X2) Execute(frame *rtda.Frame) {
	val1 := frame.PopSlot()
	val2 := frame.PopSlot()
	val3 := frame.PopSlot()
	val4 := frame.PopSlot()
	frame.PushSlot(val2)
	frame.PushSlot(val1)
	frame.PushSlot(val4)
	frame.PushSlot(val3)
	frame.PushSlot(val2)
	frame.PushSlot(val1)
}
