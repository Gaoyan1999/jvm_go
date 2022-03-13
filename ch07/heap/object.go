package heap

type Object struct {
	class  *Class
	fields Slots
}

func (ref *Object) Fields() Slots {
	return ref.fields
}

func (ref *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(ref.class)
}
