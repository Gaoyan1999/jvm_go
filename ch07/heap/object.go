package heap

type Object struct {
	Class  *Class
	fields Slots
}

func (ref *Object) Fields() Slots {
	return ref.fields
}

func (ref *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(ref.Class)
}
