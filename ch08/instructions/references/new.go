package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type New struct{ base.Index16Instruction }

func (new *New) Execute(frame *rtda.Frame) {
	cp := frame.Method.ClassMember.Class().ConstantPool()
	classRef := cp.GetConstant(new.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted {
		frame.RevertNextPC()
		base.InitClass(frame.Thread,class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack.PushRef(ref)
}
