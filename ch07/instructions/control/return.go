package control

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type Return struct{ base.NoOperandsInstruction } // Return void from method
type AReturn struct{ base.NoOperandsInstruction } // Return reference from method
type DReturn struct{ base.NoOperandsInstruction } // Return double from method
type FReturn struct{ base.NoOperandsInstruction } // Return float from method
type IReturn struct{ base.NoOperandsInstruction } // Return int from method
type LReturn struct{ base.NoOperandsInstruction } // Return long from method

func (ins *Return) Execute(frame *rtda.Frame){
	frame.Thread.PopFrame()
}

func (ins *IReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack.PopInt()
	invokerFrame.OperandStack.PushInt(retVal)
}

func (ins *AReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack.PopRef()
	invokerFrame.OperandStack.PushRef(retVal)
}

func (ins *DReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack.PopDouble()
	invokerFrame.OperandStack.PushDouble(retVal)
}

func (ins *FReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack.PopFloat()
	invokerFrame.OperandStack.PushFloat(retVal)
}

func (ins *LReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack.PopFloat()
	invokerFrame.OperandStack.PushFloat(retVal)
}