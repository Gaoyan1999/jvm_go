package constants

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }

func (ins *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(nil)
}

type DCONST_0 struct{ base.NoOperandsInstruction }

func (ins *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

func (ins *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushDouble(1.0)
}

type FCONST_0 struct{ base.NoOperandsInstruction }

func (ins *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (ins *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(1.0)
}

type FCONST_2 struct {base.NoOperandsInstruction}
func (ins *FCONST_2) Execute(frame *rtda.Frame){
	frame.OperandStack.PushFloat(2.0)
}

type ICONST_M1 struct{ base.NoOperandsInstruction }

func (ins *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (ins *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (ins *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (ins *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (ins *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (ins *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (ins *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(5)
}

type LCONST_0 struct{ base.NoOperandsInstruction }

func (ins *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (ins *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushLong(1)
}
