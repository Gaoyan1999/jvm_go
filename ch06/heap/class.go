package heap

import (
	. "jvmgo/ch05/classfile"
	"jvmgo/ch06/classfile"
)

type Class struct {
	accessFlags    uint16
	name           string
	superClassName string
	interfaceNames []string
	constantPool   *ConstantPool
	fields         []string
	methods        []string
	//loader *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   int
	staticVars        string
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	// TODO: pool, field and methods
	return class
}
