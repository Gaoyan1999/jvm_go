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
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUnit16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)

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
	} else if cf.majorVersion >= 46 && cf.majorVersion <= 52 && cf.minorVersion == 0 {
		return
	}
	panic("NO SUPPORT: Version other than JDK 8 are not supported at present.")
}
