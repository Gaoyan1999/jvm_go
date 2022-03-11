package heap

func (class *Class) isAssignableFrom(other *Class) bool {
	s, t := other, class
	if s == t {
		return true
	}
	if !t.IsInterface() {
		// s 是 t 的子类
		return s.isSubClassOf(other)
	} else {
		// t 是interface, s 是它的实现
		return s.isImplements(t)
	}
}

func (class *Class) isSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}
func (class *Class) isImplements(iFace *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iFace || i.isSubInterfaceOf(iFace) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSubInterfaceOf(iFace *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == iFace || superInterface.isSubInterfaceOf(iFace) {
			return true
		}
	}
	return false
}
