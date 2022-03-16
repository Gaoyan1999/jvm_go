package heap

import "jvmgo/ch07/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (ref *MemberRef) copyMemberInfo(refInfo *classfile.ConstantMemberRefInfo) {
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
}
func (ref *MemberRef) Name() string {
	return ref.name
}
func (ref *MemberRef) Descriptor() string {
	return ref.descriptor
}
