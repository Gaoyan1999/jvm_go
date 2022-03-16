package base

import (
	"fmt"
	"jvmgo/ch08/heap"
	"jvmgo/ch08/rtda"
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

	// Hack
	if method.IsNative() {
		if method.Name() =="registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name, method.Name(), method.Descriptor()))
		}
	}
}
