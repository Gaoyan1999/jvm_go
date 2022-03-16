package heap

import "jvmgo/ch08/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (fieldRef *FieldRef) ResolvedField() *Field {
	if fieldRef.field == nil {
		fieldRef.resolveFieldRef()
	}
	return fieldRef.field
}

func (fieldRef *FieldRef) resolveFieldRef() {
	d := fieldRef.cp.class
	targetClass := fieldRef.ResolvedClass()
	field := lookupField(targetClass, fieldRef.name, fieldRef.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	fieldRef.field = field
}

func lookupField(c *Class, fieldName, descriptor string) *Field {
	// 1. 在 C 的 field 中查找
	for _, field := range c.fields {
		if field.name == fieldName && field.descriptor == descriptor {
			return field
		}
	}
	// 2. 在 C 的 interface 中递归查找
	for _, field := range c.interfaces {
		if field := lookupField(field, fieldName, descriptor); field != nil {
			return field
		}
	}

	// 3. 在 C 的 super class 中递归查找
	if c.superClass != nil {
		return lookupField(c.superClass, fieldName, descriptor)
	}

	return nil
}
