package classfile
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	count:= int(reader.readUint16())
	pool:=make([]ConstantInfo,count)
	// len: count-1
	for i:=1;i<count-1;i++ {
		pool[i] = readConstantInfo(reader,pool)
		// TODO: handle long info and float
	}
	return pool
}

func  (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info:= cp[index]; info != nil {
		return info
	}
	panic("Invalid constant info.")
}

//func (cp ConstantPool)getNameAndType(index uint16) (string,string){
//	return "", ""
//}
//func getClassName(index uint16) string  {
//	return ""
//}

func (pool ConstantPool) getUtf8(index uint16) string {
	pool.getConstantInfo(index)
	return "mock_utf-8"

}



func (cp *ConstantPool) getClassName(index uint16) string  {
	return "mock-class-name"
}

