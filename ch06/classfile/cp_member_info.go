package classfile

/*
CONSTANT_Fieldref_info {
   u1 tag;
   u2 class_index;
   u2 name_and_type_index;
}
*/
type ConstantMemberRefInfo struct {
	ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (memberInfo *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	memberInfo.classIndex = reader.readUint16()
	memberInfo.nameAndTypeIndex = reader.readUint16()
}

func (memberInfo *ConstantMemberRefInfo) ClassName() string {
	return memberInfo.getClassName(memberInfo.classIndex)
}

func (memberInfo *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return memberInfo.getNameAndType(memberInfo.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct{ ConstantMemberRefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberRefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberRefInfo }
