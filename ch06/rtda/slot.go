package rtda

import "jvmgo/ch06/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
