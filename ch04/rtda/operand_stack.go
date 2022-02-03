package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].num = val
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	stack.size--
	return stack.slots[stack.size].num
}

func (stack *OperandStack) PushFloat(val float32) {
	stack.slots[stack.size].num = int32(math.Float32bits(val))
	stack.size++
}

func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	return math.Float32frombits(uint32(stack.slots[stack.size].num))
}

func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size].num = int32(val)
	stack.slots[stack.size+1].num = int32(val >> 32)
	stack.size += 2
}
func (stack *OperandStack) PopLong() int64 {
	stack.size -= 2
	low := stack.slots[stack.size].num
	high := stack.slots[stack.size+1].num
	return int64(high)<<32 | int64(low)
}

func (stack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	stack.PushLong(int64(bits))
	stack.size += 2
}

func (stack *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(stack.PopLong()))
}