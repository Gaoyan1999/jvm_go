package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/load"
	"jvmgo/ch05/instructions/math"
	"jvmgo/ch05/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (ins *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	// TODO: complete it
	case 0x15:
		inst := &load.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		ins.modifiedInstruction = inst
	case 0x84:
		inst := &math.IInc{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
		ins.modifiedInstruction = inst
	default:
		panic("Unsupported opcode:"+  string(opcode))
	}
}
func (ins *WIDE) Execute(frame *rtda.Frame) {
	ins.modifiedInstruction.Execute(frame)
}
