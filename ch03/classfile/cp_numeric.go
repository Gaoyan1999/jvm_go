package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

// int,boolean,byte,short
// 读取一个 32 位的数据，将它转换为int
func (self *ConstantIntegerInfo)  readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// Float
// readInfo（）先读取一个uint32数据，然后调用math包的Float32frombits（）函数把它转换成float32类
type ConstantFloatInfo struct {
	val float32
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
type ConstantDoubleInfo struct {
	val float64
}

func (self * ConstantDoubleInfo) readInfo(reader *ClassReader) {
	readUint64 := reader.readUint64()
	self.val = math.Float64frombits(readUint64)
}

// Long
type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}