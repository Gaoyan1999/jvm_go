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
	refs := arrRef.Refs()
	checkIndex(len(refs),index)


	switch instr.atype {
	case heap.AT_BOOLEAN:
		frame.PushInt(int32(arrRef.GetBytes()[index]))
	case heap.AT_CHAR:
		frame.PushInt(int32(arrRef.GetChars()[index]))
	case heap.AT_SHORT:
		frame.PushInt(int32(arrRef.GetShorts()[index]))
	case heap.AT_INT:
		frame.PushInt(arrRef.GetInts()[index])
	case heap.AT_LONG:
		frame.PushLong(arrRef.GetLongs()[index])
	case heap.AT_FLOAT:
		frame.PushFloat(arrRef.GetFloats()[index])
	case heap.AT_DOUBLE:
		frame.PushDouble(arrRef.GetDoubles()[index])
	default:
		frame.PushRef(arrRef.GetRefs()[index])
	}
}




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