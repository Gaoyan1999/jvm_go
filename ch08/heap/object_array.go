package heap
 const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)
func GetPrimitiveArrayClass(loader *ClassLoader, atype uint8) *Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}

func (obj *Object) GetRefs() []*Object    { return obj.data.([]*Object) }
func (obj *Object) GetBooleans() []int8   { return obj.data.([]int8) }
func (obj *Object) GetBytes() []int8      { return obj.data.([]int8) }
func (obj *Object) GetChars() []uint16    { return obj.data.([]uint16) }
func (obj *Object) GetShorts() []int16    { return obj.data.([]int16) }
func (obj *Object) GetInts() []int32      { return obj.data.([]int32) }
func (obj *Object) GetLongs() []int64     { return obj.data.([]int64) }
func (obj *Object) GetFloats() []float32  { return obj.data.([]float32) }
func (obj *Object) GetDoubles() []float64 { return obj.data.([]float64) }
