package classfile

type AttributeInfo interface {
	readInfo(read *ClassReader)
}

func readAttributes(reader *ClassReader,cp ConstantPool)[]AttributeInfo {
	attrCount := reader.readUint16()
	attributes := make([]AttributeInfo, attrCount)
	for i:= range  attributes {
		attributes[i] = readAttribute(reader,cp)
	}
	return attributes
}
func readAttribute(reader *ClassReader,cp ConstantPool) AttributeInfo {
	index := reader.readUint16()
	attrName := cp.getUtf8(index)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32,
	cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}
	case "ConstantValue": return &ConstantValueAttribute{}
	case "Deprecated": return &DeprecatedAttribute{}
	case "Exceptions": return &ExceptionsAttribute{}
	case "LineNumberTable": return &LineNumberTableAttribute{}
	case "LocalVariableTable": return &LocalVariableTableAttribute{}
	case "SourceFile": return &SourceFileAttribute{cp: cp}
	case "Synthetic": return &SyntheticAttribute{}
	default: return &UnparsedAttribute{attrName, attrLen, nil}
	}
}