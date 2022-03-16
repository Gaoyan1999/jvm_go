package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type InvokeSpecial struct{ base.Index16Instruction }

// hack!
func (ins *InvokeSpecial) Execute(frame *rtda.Frame) {
	currentClass := frame.Method.Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(ins.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack.GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// protected方法只能被声明该方法的类或子类调用。如果违反这一规定，则抛出IllegalAccessError异常
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class != currentClass &&
		!ref.Class.IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)

}
