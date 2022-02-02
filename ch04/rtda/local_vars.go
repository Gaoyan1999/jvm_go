package rtda

import "math"

/**
A single local variable can hold a value of type boolean, byte,
char, short, int, float, reference, or returnAddress.
A pair of local variables can hold a value of type long or double.
*/
type LocalVars []Slot

func newLocalVars(maxSize uint) LocalVars {
	if maxSize > 0 {
		return make([]Slot, maxSize)
	}
	return nil
}
func (vars LocalVars) SetInt(index uint, val int32) {
	vars[index].num = val
}
func (vars LocalVars) GetInt(index uint) int32 {
	return vars[index].num
}
func (vars LocalVars) SetFloat(index uint, val float32) {
	// float32 -> uint32 -> int32
	bits := math.Float32bits(val)
	vars[index].num = int32(bits)
}
func (vars LocalVars) GetFloat(index uint) float32 {
	// int32 -> uint32 -> float32
	bits := uint32(vars[index].num)
	return math.Float32frombits(bits)
}

func (vars LocalVars) SetLong(index uint, val int64) {
	vars[index].num = int32(val)
	vars[index+1].num = int32(val >> 32)

}

func (vars LocalVars) GetLong(index uint) int64 {
	low := vars[index].num //TODO need to test
	high := vars[index+1].num
	return int64(high)<<32 | int64(low)
}

func (vars LocalVars) SetDouble(index uint, val float64) {
	// float64 -> long(int64)  -> setLong
	bits := math.Float64bits(val)
	vars.SetLong(index, int64(bits))
}
func (vars LocalVars) GetDouble(index uint) float64 {
	bits := uint64(vars.GetLong(index))
	return math.Float64frombits(bits)
}

func (vars LocalVars) SetRef(index uint, ref *Object) {
	vars[index].ref = ref
}
func (vars LocalVars) GetRef(index uint) *Object {
	return vars[index].ref
}
