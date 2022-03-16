package load

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type LLoad struct{ base.Index8Instruction }
func (ins *LLoad) Execute(frame *rtda.Frame){
	_lload(frame,ins.Index)
}

type LLoad_0 struct {base.NoOperandsInstruction}
func (ins *LLoad_0) Execute(frame *rtda.Frame)  {
	_lload(frame,0)
}

type LLoad_1 struct {base.NoOperandsInstruction}
func (ins *LLoad_1) Execute(frame *rtda.Frame)  {
	_lload(frame,1)
}

type LLoad_2 struct {base.NoOperandsInstruction}
func (ins *LLoad_2) Execute(frame *rtda.Frame)  {
	_lload(frame,2)
}

type LLoad_3 struct {base.NoOperandsInstruction}
func (ins *LLoad_3) Execute(frame *rtda.Frame)  {
	_lload(frame,3)
}

func _lload(frame *rtda.Frame,index uint){
	val := frame.LocalVars.GetLong(index)
	frame.OperandStack.PushLong(val)
}
