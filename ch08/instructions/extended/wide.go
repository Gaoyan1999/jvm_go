package extended

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/instructions/load"
	"jvmgo/ch07/instructions/math"
	"jvmgo/ch07/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (ins *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	// TODO: complete it
	case 0x15:
		inst := &load.ALoad{}
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
