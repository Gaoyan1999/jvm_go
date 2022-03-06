package heap

import "jvmgo/ch06/classfile"

// 字段和方法
type ClassMember struct {
	AccessFlags
	name       string
	descriptor string
	class      *Class
}

// copy info from class member info.
func (classMember *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	classMember.AccessFlags = AccessFlags(info.AccessFlags)
	classMember.name = info.Name()
	classMember.descriptor = info.Descriptor()
}
