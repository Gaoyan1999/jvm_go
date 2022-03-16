package heap

import "jvmgo/ch08/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
func (ref *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if ref.method == nil {
		ref.resolveInterfaceMethodRef()
	}
	return ref.method
}

func (ref *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := ref.cp.class
	class := ref.ResolvedClass()
	if !class.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(class, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.method = method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	// 如果能在接口中找到方法，就返回找到的方法，否则调用lookupMethodInInterfaces（）函数在超接口中寻找。
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
