package control

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type TableSwitch struct {
	defaultOffset int32
	low int32
	high int32
	jumpOffsets []int32
}

func (ins *TableSwitch) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()
	ins.defaultOffset =  reader.ReadInt32()
	ins.low = reader.ReadInt32()
	ins.high = reader.ReadInt32()
	jumpOffsetsCount := ins.high - ins.low +1
	ins.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (ins *TableSwitch) Execute(frame *rtda.Frame){
	idx:=frame.OperandStack.PopInt()
	var offset int
	if idx >=ins.low &&  idx <=ins.high {
		offset =  int(ins.jumpOffsets[idx-ins.low])
	}else {
		offset = int(ins.defaultOffset)
	}
	base.Branch(frame,offset)
}