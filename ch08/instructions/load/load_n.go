package load

import (
	"jvmgo/ch08/instructions/base"
)

func NewLoadN(n uint, d bool) *LoadN {
	return &LoadN{n: n, d: d}
}

// xload_n: Load XXX from local variable
type LoadN struct {
	base.NoOperandsInstruction
	n uint
	d bool // long or double
}

//func (instr *LoadN) Execute(frame *rtda.Frame) {
//	frame.Load(instr.n, instr.d)
//}
