package classfile
/**
	https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4

	CONSTANT_Class_info {
		u1 tag;
		u2 name_index;
	}
 */

type ConstantClassInfo struct {
	nameIndex uint16
	cp ConstantPool
}

func (class *ConstantClassInfo) readInfo(reader *ClassReader) {
	class.nameIndex = reader.readUint16()
}
func (class *ConstantClassInfo) name() string {
	return class.cp.getUtf8(class.nameIndex)
}

