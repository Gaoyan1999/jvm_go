package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (symRef *SymRef) ResolvedClass() *Class {
	if symRef.class == nil {
		symRef.resolveClassRef()
	}
	return symRef.class

}

// Class and interface resolution
//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-5.html#jvms-5.4.3.1

func (symRef *SymRef) resolveClassRef() {
	d := symRef.cp.class
	c := d.loader.LoadClass(symRef.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	symRef.class = c
}
