package instructions

import (
	"fmt"
	"jvmgo/ch05/instructions/base"
)
import . "jvmgo/ch05/instructions/const"
import . "jvmgo/ch05/instructions/load"
import . "jvmgo/ch05/instructions/store"

var (
	nop         = &NOP{}
	aconst_null = &ACONST_NULL{}
	iconst_m1   = &ICONST_M1{}
	iconst_0    = &ICONST_0{}
	iconst_1    = &ICONST_1{}
	iconst_2    = &ICONST_2{}
	iconst_3    = &ICONST_3{}
	iconst_4    = &ICONST_4{}
	iconst_5    = &ICONST_5{}
	lconst_0    = &LCONST_0{}
	lconst_1    = &LCONST_1{}
	fconst_0    = &FCONST_0{}
	fconst_1    = &FCONST_1{}
	fconst_2    = &FCONST_2{}
	dconst_0    = &DCONST_0{}
	dconst_1    = &DCONST_1{}
	iload_0     = &ILOAD_0{}
	iload_1     = &ILOAD_1{}
	iload_2     = &ILOAD_2{}
	iload_3     = &ILOAD_3{}
	lload_0     = &LLoad_0{}
	lload_1     = &LLoad_1{}
	lload_2     = &LLoad_2{}
	lload_3     = &LLoad_3{}
	// TODO: fLoad dLoad aLoad
	istore_0 = &ISTORE_0{}
	istore_1 = &ISTORE_1{}
	istore_2 = &ISTORE_2{}
	istore_3 = &ISTORE_3{}
	lstore_0 = &LSTORE_0{}
	lstore_1 = &LSTORE_1{}
	lstore_2 = &LSTORE_2{}
	lstore_3 = &LSTORE_3{}
	fstore_0 = &FSTORE_0{}
	fstore_1 = &FSTORE_1{}
	fstore_2 = &FSTORE_2{}
	fstore_3 = &FSTORE_3{}
)

func NewInstruction(opcode byte) base.Instruction {

	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return &BIPUSH{}
	case 0x11:
		return &SIPUSH{}
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3

	// TODO:
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3

	default:
		panic(fmt.Errorf("Unspported opcode: 0x%x!", opcode))
	}

}
