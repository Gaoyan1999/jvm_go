package classfile

import (
	"fmt"
)

/**
https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.1
*/
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []MemberInfo // ues *
	methods      []MemberInfo
	attributes   []AttributeInfo
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}
func (cf *ClassFile) superClassName() string {
	return cf.constantPool.getClassName(cf.superClass)
}
func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, indexInPool := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(indexInPool)
	}
	return interfaceNames
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	reader := new(ClassReader)
	*reader = ClassReader{data: classData}
	file := &ClassFile{}
	file.read(reader)
	return file, nil

}

func (cf *ClassFile) read(reader *ClassReader) {
	// parse class file
	//	self.readAndCheckMagic(reader) // 见3.2.3
	//	self.readAndCheckVersion(reader) // 见3.2.4
	//	self.constantPool = readConstantPool(reader) // 见3.3
	//	self.accessFlags = reader.readUint16()
	//	self.thisClass = reader.readUint16()
	//	self.superClass = reader.readUint16()
	//	self.interfaces = reader.readUint16s()
	//	self.fields = readMembers(reader, self.constantPool) // 见3.2.8
	//	self.methods = readMembers(reader, self.constantPool)
	//	self.attributes = readAttributes(reader, self.constantPool) //见3.4

}
