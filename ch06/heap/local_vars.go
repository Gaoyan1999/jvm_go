package heap

import "math"

/**
A single local variable can hold a value of type boolean, byte,
char, short, int, float, reference, or returnAddress.
A pair of local variables can hold a value of type long or double.
*/
type LocalVars []Slot

func newSlots(maxSize uint) Slots {
	if maxSize > 0 {
		return make([]Slot, maxSize)
	}
	return nil
}
func (vars Slots) SetInt(index uint, val int32) {
	vars[index].num = val
}
func (vars Slots) GetInt(index uint) int32 {
	return vars[index].num
}
func (vars Slots) SetFloat(index uint, val float32) {
	// float32 -> uint32 -> int32
	bits := math.Float32bits(val)
	vars[index].num = int32(bits)
}
func (vars Slots) GetFloat(index uint) float32 {
	// int32 -> uint32 -> float32
	bits := uint32(vars[index].num)
	return math.Float32frombits(bits)
}

func (vars Slots) SetLong(index uint, val int64) {
	vars[index].num = int32(val)
	vars[index+1].num = int32(val >> 32)

}

func (vars Slots) GetLong(index uint) int64 {
	// 这里先转换成了 uint32 是为了解决溢出问题
	// 如果set long 的后八位为八个1，被存入到int中也就是low变量，直接将low变量变为int64，会将第一位视作+-信息
	// 所以要先将其转换为uint32类型
	low := uint32(vars[index].num)
	high := uint32(vars[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (vars Slots) SetDouble(index uint, val float64) {
	// float64 -> long(int64)  -> setLong
	bits := math.Float64bits(val)
	vars.SetLong(index, int64(bits))
}
func (vars Slots) GetDouble(index uint) float64 {
	bits := uint64(vars.GetLong(index))
	return math.Float64frombits(bits)
}

func (vars Slots) SetRef(index uint, ref *Object) {
	vars[index].ref = ref
}
func (vars Slots) GetRef(index uint) *Object {
	return vars[index].ref
}

func  (lv Slots)GetLocalVar(index uint) Slot  {
	return lv[index]
}
