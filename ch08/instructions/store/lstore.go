package store

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

// Store long into local variable
type LSTORE struct{ base.Index8Instruction }

func (ins *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, ins.Index)
}

type LSTORE_0 struct{ base.NoOperandsInstruction }

func (ins *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }

func (ins *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }

func (ins *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }

func (ins *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack.PopLong()
	frame.LocalVars.SetLong(index, val)
}