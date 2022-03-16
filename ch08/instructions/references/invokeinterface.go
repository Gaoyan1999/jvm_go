package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type InvokeInterface struct {
	index uint
}

func (ins *InvokeInterface) FetchOperands(reader *base.BytecodeReader) {
	ins.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (ins *InvokeInterface) Execute(frame *rtda.Frame) {
	pool := frame.Method.Class().ConstantPool()
	methodRef := pool.GetConstant(ins.index).(*heap.InterfaceMethodRef)
	resolvedInterfaceMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedInterfaceMethod.IsStatic() || resolvedInterfaceMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack.GetRefFromTop(resolvedInterfaceMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class.IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class, methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
