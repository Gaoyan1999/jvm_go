package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type CheckCast struct {base.Index16Instruction}

func (ins *CheckCast) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}
	cp := frame.Method.Class().ConstantPool()
	classRef := cp.GetConstant(ins.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
