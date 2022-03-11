package references

import (
	"jvmgo/ch06/heap"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type GetField struct {base.Index16Instruction}

func (ins *GetField) Execute(frame *rtda.Frame){
	pool := frame.Method.Class().ConstantPool()
	fieldRef := pool.GetConstant(ins.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	stack := frame.OperandStack
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': stack.PushInt(slots.GetInt(slotId))
	case 'F': stack.PushFloat(slots.GetFloat(slotId))
	case 'J': stack.PushLong(slots.GetLong(slotId))
	case 'D': stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[': stack.PushRef(slots.GetRef(slotId))
	}
}
