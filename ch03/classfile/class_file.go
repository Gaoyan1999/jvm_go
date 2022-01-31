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
	fields       []MemberInfo // ues * in book
	methods      []MemberInfo // ues * in book
	attributes   []AttributeInfo
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}
func (cf *ClassFile) superClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
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
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	//	self.constantPool = readConstantPool(reader) // 见3.3
	//TODO: parse access flag.

	//	self.accessFlags = reader.readUint16()
	//	self.thisClass = reader.readUint16()
	//	self.superClass = reader.readUint16()
	//	self.interfaces = reader.readUint16s()
	//	self.fields = readMembers(reader, self.constantPool) // 见3.2.8
	//	self.methods = readMembers(reader, self.constantPool)
	//	self.attributes = readAttributes(reader, self.constantPool) //见3.4

}
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magicNumber := reader.readUint32()
	if magicNumber != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic number error.")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.majorVersion = reader.readUint16()
	cf.minorVersion = reader.readUint16()
	// support JDK 8  major version [45 - 52]
	if cf.majorVersion == 45 {
		return
	} else if cf.majorVersion >= 46 && cf.majorVersion <=52 && cf.minorVersion == 0 {
		return
	}
	panic("NO SUPPORT: Version other than JDK 8 are not supported at present.")
}
