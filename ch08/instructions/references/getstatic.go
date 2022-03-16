package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

// 取出类的静态变量，推入栈顶
type GetStatic struct{ base.Index16Instruction }

func (getStatic *GetStatic) Execute(frame *rtda.Frame) {
	constantPool := frame.Method.Class().ConstantPool()
	fieldRef := constantPool.GetConstant(getStatic.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := field.Class()
	slotId := field.SlotId()
	slots := class.StaticFieldSlots()
	descriptor := field.Descriptor()
	stack := frame.OperandStack
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case '[', 'L':
		stack.PushRef(slots.GetRef(slotId))
	}
}
