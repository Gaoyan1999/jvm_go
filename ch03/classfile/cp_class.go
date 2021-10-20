package classfile

type ConstantClassInfo struct {
	cp            ConstantPool
	// 和CONSTANT_String_info类似，name_index是常量池索引，指向CONSTANT_Utf8_info常量
	nameIndex     uint16
}
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
