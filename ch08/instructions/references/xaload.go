package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type XALoad struct {
	base.NoOperandsInstruction
	atype byte
}

func (instr *XALoad) Execute(frame *rtda.Frame) {
	index := frame.PopInt()
	arrRef := frame.PopRef()
	checkNotNil(arrRef)
	//checkIndex(len(arrRef),index)
	switch instr.atype {
	case heap.AT_BOOLEAN:
		frame.PushInt(int32(arrRef.Bytes()[index]))
	case heap.AT_CHAR:
		frame.PushInt(int32(arrRef.Chars()[index]))
	case heap.AT_SHORT:
		frame.PushInt(int32(arrRef.Shorts()[index]))
	case heap.AT_INT:
		frame.PushInt(arrRef.Ints()[index])
	case heap.AT_LONG:
		frame.PushLong(arrRef.Longs()[index])
	case heap.AT_FLOAT:
		frame.PushFloat(arrRef.Floats()[index])
	case heap.AT_DOUBLE:
		frame.PushDouble(arrRef.Doubles()[index])
	default:
		frame.PushRef(arrRef.Refs()[index])
	}
}
func NewIALoad() *XALoad { return &XALoad{atype: heap.AT_INT} }
func NewLALoad() *XALoad { return &XALoad{atype: heap.AT_LONG} }
func NewFALoad() *XALoad { return &XALoad{atype: heap.AT_FLOAT} }
func NewDALoad() *XALoad { return &XALoad{atype: heap.AT_DOUBLE} }
func NewAALoad() *XALoad { return &XALoad{atype: 0} }
func NewBALoad() *XALoad { return &XALoad{atype: heap.AT_BYTE} }
func NewCALoad() *XALoad { return &XALoad{atype: heap.AT_CHAR} }
func NewSALoad() *XALoad { return &XALoad{atype: heap.AT_SHORT} }



type AALoad struct{ base.NoOperandsInstruction }
type BALoad struct{ base.NoOperandsInstruction }
type CALoad struct{ base.NoOperandsInstruction }
type DALoad struct{ base.NoOperandsInstruction }
type FALoad struct{ base.NoOperandsInstruction }
type IALoad struct{ base.NoOperandsInstruction }
type LALoad struct{ base.NoOperandsInstruction }
type SALoad struct{ base.NoOperandsInstruction }

func (aALoad *AALoad) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack
	idx := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs),idx)
	stack.PushRef(refs[idx])
}


func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int ,index int32){
	if index < 0  || index>=int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}