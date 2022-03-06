package classfile

import "math"

// Integer
type ConstantIntegerInfo struct {
	value int32
}

func (info *ConstantIntegerInfo) Value() int32 {
	return info.value
}

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	ci.value = int32(reader.readUint32())
}

// Float
type ConstantFloatInfo struct {
	value float32
}
func (cf *ConstantFloatInfo) Value() float32 {
	return cf.value
}

func (cf *ConstantFloatInfo) readInfo(reader *ClassReader) {
	cf.value = math.Float32frombits(reader.readUint32())
}

// Long
type ConstantLongInfo struct {
	value int64
}
func (cl *ConstantLongInfo) Value() int64{
	return cl.value
}

func (cl *ConstantLongInfo) readInfo(reader *ClassReader) {
	cl.value = int64(reader.readUint64())
}

// Double
type ConstantDoubleInfo struct {
	value float64
}

func (cd *ConstantDoubleInfo) Value() float64  {
	return  cd.value
}

func (cd *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	cd.value = math.Float64frombits(reader.readUint64())
}
