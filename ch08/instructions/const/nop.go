package constants

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (ins *NOP) Execute(frame *rtda.Frame) {
}