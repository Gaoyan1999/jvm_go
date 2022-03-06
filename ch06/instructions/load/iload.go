package load

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// The index is an unsigned byte that must be an index into the local variable array of the current frame
type ILOAD struct{ base.Index8Instruction }

func (ins *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, ins.Index)
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (ins *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (ins *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (ins *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (ins *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars.GetInt(index)
	frame.OperandStack.PushInt(val)
}
