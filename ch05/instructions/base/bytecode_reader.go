package base

type BytecodeReader struct {
	code []byte
	PC   int
}

func (reader *BytecodeReader) Reset(code []byte, pc int) {
	reader.code = code
	reader.PC = pc

}
func (reader *BytecodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}
func (reader *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(reader.ReadUint8())
	byte2 := uint16(reader.ReadUint8())
	return (byte1 << 8) | byte2
}
func (reader *BytecodeReader) ReadUint8() uint8 {
	i := reader.code[reader.PC]
	reader.PC++
	return i
}
func (reader *BytecodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}
func (reader *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(reader.ReadUint8())
	byte2 := int32(reader.ReadUint8())
	byte3 := int32(reader.ReadUint8())
	byte4 := int32(reader.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
func (reader *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints:=make([]int32,n)
	for i:=range  ints {
		ints[i] = reader.ReadInt32()
	}
	return ints
}
func (reader *BytecodeReader) SkipPadding(){
	for reader.PC%4 !=0 {
		reader.ReadUint8()
	}
}
