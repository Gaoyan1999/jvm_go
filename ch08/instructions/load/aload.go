package load

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

// The index is an unsigned byte that must be an index into the local variable array of the current frame
type ALoad struct{ base.Index8Instruction }

func (ins *ALoad) Execute(frame *rtda.Frame) {
	_aload(frame, ins.Index)
}

type ALoad_0 struct{ base.NoOperandsInstruction }

func (ins *ALoad_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALoad_1 struct{ base.NoOperandsInstruction }

func (ins *ALoad_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALoad_2 struct{ base.NoOperandsInstruction }

func (ins *ALoad_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALoad_3 struct{ base.NoOperandsInstruction }

func (ins *ALoad_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars.GetRef(index)
	frame.OperandStack.PushRef(val)
}
