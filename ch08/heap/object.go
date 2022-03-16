package heap

type Object struct {
	Class  *Class
	data interface{}
}

func (ref *Object) Fields() Slots {
	return ref.data.(Slots)
}

func (ref *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(ref.Class)
}
