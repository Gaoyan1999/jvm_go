package classfile

import "fmt"

// Java虚拟机规范定义的class文件格式
type ClassFile struct {
	// magic uint32
	minorVersion uint16 //char
	majorVersion uint16 //char
	constantPool ConstantPool
	accessFlags uint16
	thisClass 	uint16
	superClass  uint16
	interfaces	[]uint16
	fields 		[]*MemberInfo
	methods 	[]*MemberInfo
	attributes  []AttributeInfo
}

// 将二进制的data 转换为结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func(){
		if r:=recover(); r!=nil {
			var ok bool
			 err,ok = r.(error)
			 if !ok {
				 err = fmt.Errorf("%v",r)
			 }
		}
	}()
	cr := &ClassReader{data: classData}
	cf  = &ClassFile{}
	cf.read(cr)
	return
}
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader) // 见3.2.3
	self.readAndCheckVersion(reader) // 见3.2.4
	self.constantPool = readConstantPool(reader) // 见3.3
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool) // 见3.2.8
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool) //见3.4
}
// 检查魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	//读取4个字节,8个16进制数 
	magic := reader.readUint32() // u4
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16() //u2
	self.majorVersion = reader.readUint16() //u2
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 { // 只能为0
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
func (self *ClassFile) MinorVersion() uint16 {...} // getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
} // getter
func (self *ClassFile) ConstantPool() ConstantPool {...} // getter
func (self *ClassFile) AccessFlags() uint16 {...} // getter
func (self *ClassFile) Fields() []*MemberInfo {...} // getter
func (self *ClassFile) Methods() []*MemberInfo {...} // getter
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClass(self.thisClass)
}
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return  ""
}
// 从 constant pool 中找接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i,cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}