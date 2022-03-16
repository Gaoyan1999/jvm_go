package constants

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
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
