package heap

import "jvmgo/ch07/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = uint(codeAttr.MaxStack)
		m.maxLocals = uint(codeAttr.MaxLocals)
		m.code = codeAttr.Code
	}
}
func (m *Method) calcArgSlotCount() {
	parsedDescriptor :=parseMethodDescriptor(m.descriptor)
	for _,paraType := range parsedDescriptor.parameterTypes {
		m.argSlotCount++
		if paraType == "J" || paraType == "D" {
			m.argSlotCount ++
		}
	}
	if !m.IsStatic() {
		m.argSlotCount ++
	}
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}
func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}
