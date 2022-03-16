package references

import (
	"jvmgo/ch07/heap"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

// 给类的某个静态变量赋值
type PutStatic struct{ base.Index16Instruction }

func (putStatic *PutStatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method
	currentClass := currentMethod.Class()
	constantPool := currentClass.ConstantPool()
	field := constantPool.GetConstant(putStatic.Index).(*heap.FieldRef).ResolvedField()
	class := field.Class()
	// 如果解析后的字段不是static throw error
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果字段为final类型，则实际操作的是静态变量，只能在初始化时赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor:=field.Descriptor()
	slotId:=field.SlotId()
	slots:= class.StaticFieldSlots()
	stack := frame.OperandStack
	switch  descriptor[0] {
	case 'Z','B','C','S','I': slots.SetInt(slotId,stack.PopInt())
	case 'F':  slots.SetFloat(slotId,stack.PopFloat())
	case 'J': slots.SetLong(slotId,stack.PopLong())
	case 'D': slots.SetDouble(slotId,stack.PopDouble())
	case '[','L': slots.SetRef(slotId,stack.PopRef())
	}
}
