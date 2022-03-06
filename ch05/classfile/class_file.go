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
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}
func (cf *ClassFile) SuperClassName() string {
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
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
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
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, nil

}

func (cf *ClassFile) read(reader *ClassReader) {
	// parse class file
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
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
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	// support JDK 8  major version [45 - 52]
	if cf.majorVersion == 45 {
		return
	} else if cf.majorVersion >= 46 && cf.majorVersion <= 52 && cf.minorVersion == 0 {
		return
	}
	panic("NO SUPPORT: Version other than JDK 8 are not supported at present.")
}

func PrintClassInfo(cf *ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.constantPool))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.fields))
	for _, f := range cf.fields {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.methods))
	for _, m := range cf.methods {
		fmt.Printf("  %s\n", m.Name())
	}
}
