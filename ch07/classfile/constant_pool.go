package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUint16())
	pool := make([]ConstantInfo, count)
	// len: count-1
	for i := 1; i < count; i++ {
		pool[i] = readConstantInfo(reader, pool)
		switch pool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return pool
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := cp[index]; info != nil {
		return info
	}
	panic("Invalid constant info.")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndTypeInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	return cp.getUtf8(nameAndTypeInfo.nameIndex), cp.getUtf8(nameAndTypeInfo.descriptorIndex)
}

func (pool ConstantPool) getUtf8(index uint16) string {
	utf8Info := pool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return classInfo.Name()
}
