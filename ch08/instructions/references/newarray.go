package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type NewArray struct {
	atype uint8
}

func (ins *NewArray) FetchOperands(reader *base.BytecodeReader) {
	ins.atype = reader.ReadUint8()
}

func (ins *NewArray) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method.Class().ClassLoader()
	arrClass := heap.GetPrimitiveArrayClass(classLoader, ins.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)

}
