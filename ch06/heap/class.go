package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

type Class struct {
	AccessFlags
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticFieldSlots  Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.AccessFlags = AccessFlags(cf.AccessFlags())
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (class *Class) isAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) StaticFieldSlots() Slots {
	return class.staticFieldSlots
}

func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (class *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}
