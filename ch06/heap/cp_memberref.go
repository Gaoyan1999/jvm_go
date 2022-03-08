package heap

import "jvmgo/ch06/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (ref *MemberRef) copyMemberInfo(refInfo *classfile.ConstantMemberRefInfo) {
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
}
