package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (ins *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, ins.Offset)
}
type GotoW struct{ offset int }

func (ins *GotoW) FetchOperands(reader *base.BytecodeReader)  {
	ins.offset = int(reader.ReadInt32())
}

func (ins *GotoW) Execute(frame *rtda.Frame) {
	base.Branch(frame, ins.offset)
}