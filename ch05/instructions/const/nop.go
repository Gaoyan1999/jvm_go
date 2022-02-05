package constants

import (
	"jvmgo/ch04/rtda"
	"jvmgo/ch05/instructions/base"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (ins *NOP) Execute(_ *rtda.Frame) {
}
