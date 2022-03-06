package classfile

type ConstantUtf8Info struct {
	str string
}

func (utf8 *ConstantUtf8Info) readInfo(reader *ClassReader) {
	len := reader.readUint16()
	bytes := reader.readBytes(uint32(len))
	utf8.str = string(bytes)
}
