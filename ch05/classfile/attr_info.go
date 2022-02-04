package classfile

const (
	Code               = "Code"
	ConstantValue      = "ConstantValue"
	Deprecated         = "Deprecated"
	Exceptions         = "Exceptions"
	LineNumberTable    = "LineNumberTable"
	LocalVariableTable = "LocalVariableTable"
	SourceFile         = "SourceFile"
	Synthetic          = "Synthetic"
)

/**
attribute_info {
   u2 attribute_name_index;
   u4 attribute_length;
   u1 info[attribute_length];
}
*/

type AttributeInfo interface {
	read(reader *ClassReader)
}

func readAttributes(reader *ClassReader, pool ConstantPool) []AttributeInfo {
	count := reader.readUint16()
	infos := make([]AttributeInfo, count)
	for i := range infos {
		infos[i] = readAttribute(reader, pool)
	}
	return infos
}
func readAttribute(reader *ClassReader, pool ConstantPool) AttributeInfo {
	nameIndex := reader.readUint16()
	name := pool.getUtf8(nameIndex)
	len := reader.readUint32()
	attrInfo := newAttributeInfo(name, len, pool)
	attrInfo.read(reader)
	return attrInfo
}

func newAttributeInfo(name string, len uint32, cp ConstantPool) AttributeInfo {
	switch name {
	case Code:
		return &CodeAttribute{ConstantPool: cp}
	case ConstantValue:
		return &ConstantValueAttribute{}
	case Deprecated:
		return &DeprecatedAttribute{}
	case Synthetic:
		return &SyntheticAttribute{}
	case SourceFile:
		return &SourceFileAttribute{ConstantPool: cp}
	case Exceptions:
		return &ExceptionsAttribute{}
	case LineNumberTable:
		return &LineNumberTableAttribute{}
	case LocalVariableTable:
		return &LocalVariableAttribute{}

	default:
		return &UnparsedAttribute{name, len, nil}

	}

}
