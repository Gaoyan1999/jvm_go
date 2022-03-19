package store

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type AAStore struct{ base.NoOperandsInstruction }
type BAStore struct{ base.NoOperandsInstruction }
type CAStore struct{ base.NoOperandsInstruction }
type DAStore struct{ base.NoOperandsInstruction }
type FAStore struct{ base.NoOperandsInstruction }
type IAStore struct{ base.NoOperandsInstruction }
type LAStore struct{ base.NoOperandsInstruction }
type SAStore struct{ base.NoOperandsInstruction }

func (ins *IAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = value
}
func (ins *AAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = value
}
func (ins *BAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(value)
}
func (ins *CAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(value)
}
func (ins *FAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	floats[index] = value
}
func (ins *DAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = value
}
func (ins *LAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	longs[index] = value
}
func (ins *SAStore) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.Shorts()
	checkIndex(len(longs), index)
	longs[index] = int16(value)
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
