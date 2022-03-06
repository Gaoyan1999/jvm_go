package comparisions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

func NewLCMP() *Lcmp  { return &Lcmp{} }
// 当两个float变量中至少有一个是NaN时，用fcmpg指令比较的结果是1，而用fcmpl指令比较的结果是-1。
func NewFCMPG() *Fcmp { return &Fcmp{g: true} }
func NewFCMPL() *Fcmp { return &Fcmp{g: false} }
func NewDCMPG() *Dcmp { return &Dcmp{g: true} }
func NewDCMPL() *Dcmp { return &Dcmp{g: false} }

type Lcmp struct{ base.NoOperandsInstruction }

func (ins *Lcmp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		frame.PushInt(1)
	} else if v1 == v2 {
		frame.PushInt(0)
	} else {
		frame.PushInt(-1)
	}
}

type Fcmp struct {
	base.NoOperandsInstruction
	// 是否可能无法比较
	g bool
}
func (ins *Fcmp) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	switch  {
	case v1>v2:
		frame.PushInt(1)
	case v1 == v2:
		frame.PushInt(0)
	case v1 < v2:
		frame.PushInt(-1)
	case ins.g:
		frame.PushInt(1)
	default:
		frame.PushInt(-1)
	}
}
// Compare double
type Dcmp struct {
	base.NoOperandsInstruction
	g bool
}

func (instr *Dcmp) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	switch {
	case v1 > v2:
		frame.PushInt(1)
	case v1 == v2:
		frame.PushInt(0)
	case v1 < v2:
		frame.PushInt(-1)
	case instr.g:
		frame.PushInt(1)
	default:
		frame.PushInt(-1)
	}
}
