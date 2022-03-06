package classfile

type ConstantStringInfo struct {
	ConstantPool
	index uint16
}

func (cs *ConstantStringInfo) readInfo(reader *ClassReader) {
	cs.index = reader.readUint16()
}

func (cs *ConstantStringInfo) getValue() string {
	return cs.getUtf8(cs.index)
}
