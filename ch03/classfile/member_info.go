package classfile

type MemberInfo struct {
	cp ConstantPool //保存常量池指针
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

func readMembers(reader *ClassReader,cp ConstantPool)[]*MemberInfo  {
	 memberCount :=reader.readUint16()
	 members :=make([]*MemberInfo,memberCount)
	 for i:= range members {
	 	members[i] = readMember(reader,cp)
	 }
	 return  members

}
func readMember(reader *ClassReader,cp ConstantPool)*MemberInfo  {
	return  &MemberInfo{
		cp: cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader,cp),

	}
}
func readAttributes(reader *ClassReader,cp ConstantPool)[]AttributeInfo{
}


func (self *MemberInfo)AccessFlags() uint16 {
	return self.accessFlags
}
// 从 constant pool 中查找name 和description
func (self *MemberInfo)Name() string{
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo)Descriptor() string  {
	return self.cp.getUtf8(self.descriptorIndex)
}
