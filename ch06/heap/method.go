package heap

import "jvmgo/ch06/classfile"

type  Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
}

func newMethods(class *Class,cfMethods []*classfile.MemberInfo) []*Method  {
	methods := make([]*Method, len(cfMethods))
	for i,cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (m *Method) copyAttributes (cfMethod *classfile.MemberInfo)  {
	if codeAttr := cfMethod.CodeAttribute();codeAttr !=nil {
		m.maxStack = uint(codeAttr.MaxStack)
		m.maxLocals = uint(codeAttr.MaxLocals)
		m.code = codeAttr.Code
	}
}
