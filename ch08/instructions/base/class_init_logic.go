package base

import (
	"jvmgo/ch08/heap"
	"jvmgo/ch08/rtda"
)

func InitClass(thread *rtda.Thread,class *heap.Class){
	class.StartInit()
	scheduleClinit(thread,class)
	initSuperClass(thread,class)
}

func scheduleClinit(thread *rtda.Thread,class *heap.Class){
	clinit := class.GetClinitMethod()
	if clinit !=nil {
		newFrame:= thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
func initSuperClass(thread *rtda.Thread,class *heap.Class){
	if !class.IsInterface() {
		superClass:= class.SuperClass()
		if superClass!=nil && !superClass.InitStarted {
			InitClass(thread,superClass)
		}
	}
}