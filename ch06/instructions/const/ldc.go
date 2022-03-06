package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type LDC struct {base.Index8Instruction}

func (ins *LDC) Execute(frame *rtda.Frame) {

}

func _ldc(frame *rtda.Frame,index uint)  {
}