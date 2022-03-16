package rtda

import "jvmgo/ch07/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
