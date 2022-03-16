package stack

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (ins *SWAP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}