package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool,refInfo *classfile.ConstantFieldrefInfo)*FieldRef  {
	ref := &FieldRef{}
	ref.cp =cp
	ref.copyMemberInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
