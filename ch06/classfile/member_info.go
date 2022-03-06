package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5
type MemberInfo struct {
	cp              ConstantPool
	AccessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	Attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	count := reader.readUint16()
	members := make([]*MemberInfo, count)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		AccessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		Attributes:      readAttributes(reader, cp),
	}
}
func (member *MemberInfo) Name() string {
	return member.cp.getUtf8(member.nameIndex)
}

func (member *MemberInfo) Descriptor() string {
	return member.cp.getUtf8(member.descriptorIndex)
}
func (member *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range member.Attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}
