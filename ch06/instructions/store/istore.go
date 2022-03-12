package store

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Store int into local variable
type ISTORE struct{ base.Index8Instruction }

func (ins *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, ins.Index)
}

type ISTORE_0 struct{ base.NoOperandsInstruction }

func (ins *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct{ base.NoOperandsInstruction }

func (ins *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct{ base.NoOperandsInstruction }

func (ins *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct{ base.NoOperandsInstruction }

func (ins *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack.PopInt()
	frame.LocalVars.SetInt(index,val)
}
