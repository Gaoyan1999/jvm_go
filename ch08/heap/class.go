package heap

import (
	"jvmgo/ch08/classfile"
	"strings"
)

type Class struct {
	AccessFlags
	Name              string
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
	InitStarted       bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.AccessFlags = AccessFlags(cf.AccessFlags())
	class.Name = cf.ClassName()
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
		Class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

func (class *Class) isAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.GetPackageName() == other.GetPackageName()
}

func (class *Class) GetPackageName() string {
	if i := strings.LastIndex(class.Name, "/"); i >= 0 {
		return class.Name[:i]
	}
	return ""
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) StaticFieldSlots() Slots {
	return class.staticFieldSlots
}
func (class *Class) SuperClass() *Class {
	return class.superClass
}

func (class *Class) ClassLoader() *ClassLoader {
	return class.loader
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

func (class *Class) StartInit() {
	class.InitStarted = true
}

func (class *Class) GetClinitMethod() *Method {
	return class.getStaticMethod("<clinit>", "()V")
}

func (class *Class) ArrayClass() *Class  {
	arrayClassName := getArrayClassName(class.Name)
	return class.loader.LoadClass(arrayClassName)

}
