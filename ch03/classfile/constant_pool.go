package classfile
type ConstantPool struct {

}

func (cp *ConstantPool) getClassName(index uint16) string  {
	return "mock-class-name"
}
