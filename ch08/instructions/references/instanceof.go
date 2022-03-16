package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/heap"
	"jvmgo/ch07/rtda"
)

type InstanceOf struct{ base.Index16Instruction }

func (ins *InstanceOf) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method.Class().ConstantPool()
	classRef := cp.GetConstant(ins.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	}else {
		stack.PushInt(0)
	}
}