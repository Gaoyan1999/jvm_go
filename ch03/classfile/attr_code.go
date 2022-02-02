package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type CodeAttribute struct {
	ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}
type ExceptionTableEntry struct {
	startPc     uint16
	endPc       uint16
	handlerPc   uint16
	catchTypePc uint16
}

func (attr *CodeAttribute) read(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocals = reader.readUint16()
	codeLen := reader.readUint32()
	attr.code = reader.readBytes(codeLen)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.ConstantPool)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	len := reader.readUint16()
	table := make([]*ExceptionTableEntry, len)
	for i := range table {
		table[i] = &ExceptionTableEntry{
			startPc:     reader.readUint16(),
			endPc:       reader.readUint16(),
			handlerPc:   reader.readUint16(),
			catchTypePc: reader.readUint16(),
		}
	}
	return table
}
