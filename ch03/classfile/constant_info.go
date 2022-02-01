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

// return a type of constant info according to the tag
func newConstantInfo(tag uint8, pool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: pool}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_String:
		return &ConstantStringInfo{ConstantPool: pool}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberRefInfo{ConstantPool: pool}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberRefInfo{ConstantPool: pool}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantMethodrefInfo{ConstantMemberRefInfo{ConstantPool: pool}}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: invalid constant tag.")
	}
}
