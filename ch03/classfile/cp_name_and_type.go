package classfile
// CONSTANT_NameAndType_info给出字段或方法的名称和描述符。
//CONSTANT_Class_info和CONSTANT_NameAndType_info加在一起可以唯一确定一个字段或者方法。
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}