package references

import (
	"jvmgo/ch06/heap"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type New struct{ base.Index16Instruction }

func (new *New) Execute(frame *rtda.Frame) {
	cp := frame.Method.ClassMember.Class().ConstantPool()
	classRef := cp.GetConstant(new.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack.PushRef(ref)
}
