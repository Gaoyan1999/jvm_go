package instructions

import (
	"fmt"
	"jvmgo/ch05/instructions/base"
)
import . "jvmgo/ch05/instructions/const"

var (
	nop = &NOP{}
)

func NewInstruction(opcode byte) base.Instruction {

	switch opcode {
	case 0x00: return &NOP{}
	case 0x01: return &ACONST_NULL{}

	default:
		panic(fmt.Errorf("Unspported opcode: 0x%x!", opcode))
	}

}
