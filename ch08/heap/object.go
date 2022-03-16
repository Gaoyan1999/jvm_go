package heap

type Object struct {
	Class  *Class
	data interface{}
}

func (object *Object) Fields() Slots {
	return object.data.(Slots)
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(object.Class)
}
