package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type LDC struct {base.Index8Instruction}

func (ins *LDC) Execute(frame *rtda.Frame) {

}

func _ldc(frame *rtda.Frame,index uint)  {
}