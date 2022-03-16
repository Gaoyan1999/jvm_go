package control

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type LookupSwitch struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (ins *LookupSwitch) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ins.defaultOffset = reader.ReadInt32()
	ins.npairs = reader.ReadInt32()
	ins.matchOffsets = reader.ReadInt32s(ins.npairs * 2)
}

func (ins *LookupSwitch) Execute(frame *rtda.Frame) {
	key := frame.OperandStack.PopInt()
	for i := int32(0); i < ins.npairs*2; i += 2 {
		if key == ins.matchOffsets[i] {
			offset := ins.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(ins.defaultOffset))
}
