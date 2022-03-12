package instructions

import (
	"fmt"
	"jvmgo/ch06/instructions/base"
)
import . "jvmgo/ch06/instructions/const"
import . "jvmgo/ch06/instructions/load"
import . "jvmgo/ch06/instructions/store"
import . "jvmgo/ch06/instructions/stack"
import . "jvmgo/ch06/instructions/math"
import . "jvmgo/ch06/instructions/conversions"
import . "jvmgo/ch06/instructions/comparisions"
import . "jvmgo/ch06/instructions/control"

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
	dstore_0 = &DSTORE_0{}
	dstore_1 = &DSTORE_1{}
	dstore_2 = &DSTORE_2{}
	dstore_3 = &DSTORE_3{}
	pop      = &POP{}
	pop2     = &POP2{}
	dup      = &DUP{}
	dup_x1   = &DUP_X1{}
	dup_x2   = &DUP_X2{}
	dup2     = &DUP2{}
	dup2_x1  = &DUP2_X1{}
	dup2_x2  = &DUP2_X2{}
	swap     = &SWAP{}
	iadd     = NewIAdd()
	ladd     = NewLAdd()
	fadd     = NewFAdd()
	dadd     = NewDAdd()
	isub     = NewISub()
	lsub     = NewLSub()
	fsub     = NewFSub()
	dsub     = NewDSub()
	imul     = NewIMul()
	lmul     = NewLMul()
	fmul     = NewFMul()
	dmul     = NewDMul()
	idiv     = NewIDiv()
	ldiv     = NewLDiv()
	fdiv     = NewFDiv()
	ddiv     = NewDDiv()
	irem     = NewIRem()
	lrem     = NewLRem()
	frem     = NewFRem()
	drem     = NewDRem()
	ineg     = NewINeg()
	lneg     = NewLNeg()
	fneg     = NewFNeg()
	dneg     = NewDNeg()
	ishl     = NewIShl()
	lshl     = NewLShl()
	ishr     = NewIShr()
	lshr     = NewLShr()
	iushr    = NewIUShr()
	lushr    = NewLUShr()
	iand     = NewIAnd()
	land     = NewLAnd()
	ior      = NewIOr()
	lor      = NewLOr()
	ixor     = NewIXor()
	lxor     = NewLXor()
	i2l      = NewI2L()
	i2f      = NewI2F()
	i2d      = NewI2D()
	l2i      = NewL2I()
	l2f      = NewL2F()
	l2d      = NewL2D()
	f2i      = NewF2I()
	f2l      = NewF2L()
	f2d      = NewF2D()
	d2i      = NewD2I()
	d2l      = NewD2L()
	d2f      = NewD2F()
	i2b      = NewI2B()
	i2c      = NewI2C()
	i2s      = NewI2S()
	lcmp     = NewLCMP()
	fcmpl    = NewFCMPL()
	fcmpg    = NewFCMPG()
	dcmpl    = NewDCMPL()
	dcmpg    = NewDCMPG()
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
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	// TODO: astore 0x46 ~ 0x56
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &IInc{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return NewIfEQ()
	case 0x9a:
		return NewIfNE()
	case 0x9b:
		return NewIfLT()
	case 0x9c:
		return NewIfGE()
	case 0x9d:
		return NewIfGT()
	case 0x9e:
		return NewIfLE()
	case 0x9f:
		return NewIfICmpEQ()
	case 0xa0:
		return NewIfICmpNE()
	case 0xa1:
		return NewIfICmpLT()
	case 0xa2:
		return NewIfICmpGE()
	case 0xa3:
		return NewIfICmpGT()
	case 0xa4:
		return NewIfICmpLE()
	case 0xa5:
		return NewIfACmpEQ()
	case 0xa6:
		return NewIfACmpNE()
	case 0xa7:
		return &GOTO{}
	case 0xaa:
		return &TableSwitch{}
	case 0xab:
		return &LookupSwitch{}

	default:
		panic(fmt.Errorf("Unspported opcode: 0x%x!", opcode))
	}

}
