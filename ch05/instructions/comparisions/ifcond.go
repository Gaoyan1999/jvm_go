package comparisions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IFEQ struct {base.BranchInstruction}
type IFNE struct {base.BranchInstruction}
// x<0
type IFLT struct {base.BranchInstruction}
// x<=0
type IFLE struct {base.BranchInstruction}
// x>0
type IFGT struct {base.BranchInstruction}
// x>=0
type IFGE struct {base.BranchInstruction}

func (ins *IFEQ) Execute(frame *rtda.Frame)  {
	val:=frame.PopInt()
	if val == 0 {
		base.Branch(frame,ins.Offset)
	}


}