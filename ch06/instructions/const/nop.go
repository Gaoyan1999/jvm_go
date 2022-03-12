package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (ins *NOP) Execute(frame *rtda.Frame) {
}