package references

import (
	"fmt"
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

// Invoke instance method; dispatch based on class
type InvokeVirtual struct{ base.Index16Instruction }


func (ins *InvokeVirtual) Execute(frame *rtda.Frame) {
	currentClass := frame.Method.Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack.GetRefFromTop(method.ArgSlotCount() - 1)
	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack, methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	if method.IsProtected() && method.Class().IsSuperClassOf(currentClass) &&
		method.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class != currentClass && !ref.Class.IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class, methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V": fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V": fmt.Printf("%c\n", stack.PopInt())
	case "(B)V": fmt.Printf("%v\n", stack.PopInt())
	case "(S)V": fmt.Printf("%v\n", stack.PopInt())
	case "(I)V": fmt.Printf("%v\n", stack.PopInt())
	case "(F)V": fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V": fmt.Printf("%v\n", stack.PopLong())
	case "(D)V": fmt.Printf("%v\n", stack.PopDouble())
	default: panic("println: " + descriptor)
	}
	stack.PopRef()
}