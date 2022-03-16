package rtda

import "jvmgo/ch08/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
