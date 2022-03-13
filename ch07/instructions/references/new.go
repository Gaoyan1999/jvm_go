package references

import (
	"jvmgo/ch07/heap"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
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
