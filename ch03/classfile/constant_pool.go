package classfile
// 常量信息
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
// 1. 常量池表头的cp大小比实际大1
// 2. 若表头给出的值是n constant pool 有效索引是 [1,n-1]
// 3. CONSTANT_Long_info和CONSTANT_Double_info各占两个位置。
//也就是说，如果常量池中存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些数也会变成无效索引。
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	 cpCount:=int(reader.readUint16())
	  cp := make([]ConstantInfo,cpCount)
	  for i:=1;i<cpCount; i++ {
	  	cp[i] = readConstantInfo(reader,cp)
		  switch cp[i].(type) {
		  case *ConstantLongInfo,*ConstantDoubleInfo:
		  i++ // 占2位
		  }
	  }
	  return cp
}

func readConstantInfo(reader *ClassReader, cp []ConstantInfo) ConstantInfo {
	// 读取常量信息
	tag := reader.readUint8()
	// 创建具体的constant
	c := newConstantInfo(tag, cp)
	// 读取常量信息
	c.readInfo(reader)
	return c
}
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_Float: return &ConstantFloatInfo{}
	case CONSTANT_Long: return &ConstantLongInfo{}
	case CONSTANT_Double: return &ConstantDoubleInfo{}
	case CONSTANT_Utf8: return &ConstantUtf8Info{}
	case CONSTANT_String: return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class: return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType: return &ConstantMemberRefInfo{}
	case CONSTANT_MethodHandle: return &ConstantMemberRefInfo{}
	case CONSTANT_InvokeDynamic: return &ConstantMemberRefInfo{}
	default: panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo:= self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}
func (self ConstantPool) getClassName(index uint16) string {
	ntInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return  self.getUtf8(ntInfo.nameIndex)

}
// 从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return  info.str
}