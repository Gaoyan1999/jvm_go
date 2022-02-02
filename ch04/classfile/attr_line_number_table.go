package classfile

/**
LineNumberTable_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
   u2 line_number_table_length;
   {   u2 start_pc;
       u2 line_number;
   } line_number_table[line_number_table_length];
}
*/

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (attr *LineNumberTableAttribute) read(reader *ClassReader) {
	len := reader.readUint16()
	table := make([]*LineNumberTableEntry, len)
	for i := range table {
		table[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
