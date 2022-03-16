package references

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

// 创建引用类型的数组
type AnewArray struct {base.Index16Instruction}

func (ins *AnewArray) Execute(frame *rtda.Frame)  {
	constantPool := frame.Method.Class().ConstantPool()
	classRef:=constantPool.GetConstant(ins.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack
	count := stack.PopInt()
	if count<0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrayClass := componentClass.ArrayClass()
	arr := arrayClass.NewArray(uint(count))
	stack.PushRef(arr)
}