package classfile

/**
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/

type LocalVariableAttribute struct {
	lineNumberTable []*LocalVariableEntry
}

type LocalVariableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (attr *LocalVariableAttribute) read(reader *ClassReader) {
	len := reader.readUint16()
	table := make([]*LocalVariableEntry, len)
	for i := range table {
		table[i] = &LocalVariableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}