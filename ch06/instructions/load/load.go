package load

import (
	"jvmgo/ch05/instructions/base"
)

func NewLoad(d bool) *Load {
	return &Load{d: d}
}

// xload: Load XXX from local variable
type Load struct {
	base.Index8Instruction
	d bool // long or double
}

//func (instr *Load) Execute(frame *rtda.Frame) {
//	frame.Load(instr.Index, instr.d)
//}
