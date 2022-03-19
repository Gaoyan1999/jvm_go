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

func (object *Object)SetRefVar(name string,descriptor string,ref *Object) {
	field := object.Class.getField(name, descriptor, false)
	slots := object.data.(Slots)
	slots.SetRef(field.slotId,ref)
}

func (object *Object) GetRefVar(name, descriptor string) *Object {
	field := object.Class.getField(name, descriptor, false)
	slots := object.data.(Slots)
	return slots.GetRef(field.slotId)
}