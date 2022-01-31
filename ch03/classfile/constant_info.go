package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, pool ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	return newConstantInfo(tag, pool)
}

//  return a type of constant info according to the tag
func newConstantInfo(tag uint8, pool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: pool}
	default:
		panic("java.lang.ClassFormatError: invalid constant tag.")
	}
}
