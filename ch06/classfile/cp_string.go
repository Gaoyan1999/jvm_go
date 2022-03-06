package classfile

type ConstantStringInfo struct {
	ConstantPool
	index uint16
}

func (cs *ConstantStringInfo) readInfo(reader *ClassReader) {
	cs.index = reader.readUint16()
}

func (cs *ConstantStringInfo) GetValue() string {
	return cs.getUtf8(cs.index)
}
