package classfile

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (dynamicInfo *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	dynamicInfo.BootstrapMethodAttrIndex = reader.readUint16()
	dynamicInfo.NameAndTypeIndex = reader.readUint16()
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (method *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	method.ReferenceKind = reader.readUint8()
	method.ReferenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

func (info *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	info.DescriptorIndex = reader.readUint16()
}
