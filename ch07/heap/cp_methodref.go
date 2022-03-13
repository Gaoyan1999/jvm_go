package heap

import "jvmgo/ch07/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}
func newMethodRef (cp *ConstantPool,refInfo *classfile.ConstantMethodrefInfo) * MethodRef{
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
