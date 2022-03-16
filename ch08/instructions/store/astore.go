package store

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

// Store int into local variable
type AStore struct{ base.Index8Instruction }

func (ins *AStore) Execute(frame *rtda.Frame) {
	_astore(frame, ins.Index)
}

type AStore_0 struct{ base.NoOperandsInstruction }

func (ins *AStore_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type AStore_1 struct{ base.NoOperandsInstruction }

func (ins *AStore_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type AStore_2 struct{ base.NoOperandsInstruction }

func (ins *AStore_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type AStore_3 struct{ base.NoOperandsInstruction }

func (ins *AStore_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack.PopRef()
	frame.LocalVars.SetRef(index,val)
}
