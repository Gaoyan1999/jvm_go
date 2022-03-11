package heap

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
type LDC2_W struct{ base.Index16Instruction }

func (ldc *LDC) Execute(frame *rtda.Frame){
	_ldc(frame,ldc.Index)
}
func (ldcW *LDC_W) Execute(frame *rtda.Frame){
	_ldc(frame,ldcW.Index)
}
func (ldc2W *LDC2_W) Execute(frame *rtda.Frame){
	stack := frame.OperandStack
	pool := frame.Method.class.ConstantPool()
	constant := pool.GetConstant(ldc2W.Index)
	switch constant.(type) {
	case int64: stack.PushLong(constant.(int64))
	case float64: stack.PushDouble(constant.(float64))
	default: panic("java.lang.ClassFormatError")

	}
}

func _ldc(frame *rtda.Frame,index uint){
	stack := frame.OperandStack
	pool := frame.Method.class.constantPool
	constant := pool.GetConstant(index)
	switch constant.(type) {
	case int32:stack.PushInt(constant.(int32))
	case float32:stack.PushFloat(constant.(float32))
	default: panic("TODO: ldc")
	}
}