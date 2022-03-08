package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
	slotId uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

func (field *Field)copyAttribute(cfField *classfile.MemberInfo)  {
	if valAttr := cfField.ConstantValueAttribute();valAttr !=nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (field *Field) isLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}
