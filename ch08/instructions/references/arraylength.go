package references

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type ArrayLength struct{ base.NoOperandsInstruction }

func (ins *ArrayLength) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen :=  arrRef.ArrayLength()
	stack.PushInt(arrLen)
}