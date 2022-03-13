package constants

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (ins *NOP) Execute(frame *rtda.Frame) {
}