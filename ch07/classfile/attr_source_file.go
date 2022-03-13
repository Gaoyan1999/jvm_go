package classfile

/**
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
 */

type SourceFileAttribute struct {
	sourceFileIndex uint16
	ConstantPool
}
func (attr *SourceFileAttribute) read(reader *ClassReader){
	attr.sourceFileIndex = reader.readUint16()
}

func (attr *SourceFileAttribute) FileName() string{
	return attr.getUtf8(attr.sourceFileIndex)
}