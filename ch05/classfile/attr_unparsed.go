package classfile
type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}

func (unparsed *UnparsedAttribute) read(reader *ClassReader)  {
	unparsed.info = reader.readBytes(unparsed.length)
}
