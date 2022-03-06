package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type BIPUSH struct {
	val int8
}

func (ins *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	ins.val = reader.ReadInt8()
}
func (ins *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(ins.val)
	frame.OperandStack.PushInt(i)
}

type SIPUSH struct {
	val int16
}
func (ins *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	ins.val = reader.ReadInt16()
}
func (ins *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(ins.val)
	frame.OperandStack.PushInt(i)
}
