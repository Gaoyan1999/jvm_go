package base

import "jvmgo/ch08/rtda"

func Branch(frame *rtda.Frame,offset int){
	pc:=frame.Thread.PC()
	nextPC:=pc +offset
	frame.NextPC = nextPC
}
