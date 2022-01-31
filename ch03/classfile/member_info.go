package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5
type MemberInfo struct {
	cp              ConstantPool // TODO no point
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo // TODO no point
}

func readMembers(reader *ClassReader, cp ConstantPool) []MemberInfo {
	count := reader.readUint16()
	members := make([]MemberInfo, count)
	for i := range  members {
	members[i] = readMember(reader,cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) MemberInfo {
	return MemberInfo{
		cp: cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(),
	}
}
func (member *MemberInfo) Name()  {

}

// TODO in attri
func readAttributes() []AttributeInfo  {
	res:=make([]AttributeInfo,10)
	return res
}
