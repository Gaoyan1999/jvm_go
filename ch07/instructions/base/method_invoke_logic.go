package base

import (
	"jvmgo/ch07/heap"
	"jvmgo/ch07/rtda"
)

func InvokeMethod(frame *rtda.Frame,method *heap.Method){
	thread := frame.Thread
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argLens:= int(method.ArgSlotCount())
	if argLens > 0 {
		for i:= argLens-1 ;i>=0;i-- {
		slot :=frame.OperandStack.PopSlot()
		newFrame.LocalVars.SetLocalVar(uint(i),slot)
		}
	}
}
