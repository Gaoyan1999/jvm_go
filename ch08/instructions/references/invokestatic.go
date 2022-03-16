package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type InvokeStatic struct {base.Index16Instruction}

func (ins *InvokeStatic) Execute(frame *rtda.Frame)  {
	cp := frame.Method.Class().ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame,resolvedMethod)
}
