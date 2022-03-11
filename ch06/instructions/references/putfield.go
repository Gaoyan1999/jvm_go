package references

import (
	"jvmgo/ch06/heap"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type PutField struct{ base.Index16Instruction }

func (ins *PutField) Execute(frame *rtda.Frame) {
	method := frame.Method
	currentClass := method.Class()
	pool := currentClass.ConstantPool()
	fieldRef := pool.GetConstant(ins.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		// 只能在构造函数中初始化
		if currentClass != field.Class() || method.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	slotId := field.SlotId()
	descriptor := field.Descriptor()
	stack := frame.OperandStack
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	}
}
