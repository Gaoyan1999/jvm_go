package heap

import "jvmgo/ch07/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (ref *MethodRef) ResolvedMethod() *Method {
	if ref.method == nil {
		ref.resolveMethodRef()
	}
	return ref.method
}

func (ref *MethodRef) resolveMethodRef() {
	d := ref.cp.class
	class := ref.ResolvedClass()
	if class.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(class, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.method = method
}

func  lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
	method = lookupMethodInInterfaces(class.interfaces,name,descriptor)
	}
	return method
}
func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c:=class;c!=nil;c=c.superClass {
		for _,method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return  nil
}
func lookupMethodInInterfaces(interFace []*Class, name, descriptor string) *Method {
	for _,iface := range interFace {
		for _,method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method:=lookupMethodInInterfaces(iface.interfaces,name,descriptor)
		if method !=nil {
			return method
		}
	}
	return nil
}
