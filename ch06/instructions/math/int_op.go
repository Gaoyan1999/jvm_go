package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

func NewINeg() *INeg { return &INeg{} }
func NewIAdd() *IOp  { return &IOp{op: iadd} }
func NewISub() *IOp  { return &IOp{op: isub} }
func NewIMul() *IOp  { return &IOp{op: imul} }
func NewIDiv() *IOp  { return &IOp{op: idiv, isDiv: true} }
func NewIRem() *IOp  { return &IOp{op: irem, isDiv: true} }
func NewIAnd() *IOp  { return &IOp{op: iand} }
func NewIOr() *IOp   { return &IOp{op: ior} }
func NewIXor() *IOp  { return &IOp{op: ixor} }
func NewIShl() *IOp  { return &IOp{op: ishl} }
func NewIShr() *IOp  { return &IOp{op: ishr} }
func NewIUShr() *IOp { return &IOp{op: iushr} }

type IOp struct {
	base.NoOperandsInstruction
	op    func(a, b int32) int32
	isDiv bool
}

func iadd(a, b int32) int32 { return a + b }
func isub(a, b int32) int32 { return a - b }
func imul(a, b int32) int32 { return a * b }
func idiv(a, b int32) int32 { return a / b }
func irem(a, b int32) int32 { return a % b }
func iand(a, b int32) int32 { return a & b }
func ior(a, b int32) int32  { return a | b }
func ixor(a, b int32) int32 { return a ^ b }

// 左移
func ishl(a, b int32) int32 {
	// 0x1f: 31 最多左移31位
	return a << (b & 0x1f)
}

// 右移
func ishr(a, b int32) int32 { return a >> (b & 0x1f) }

// 无符号右移
func iushr(a, b int32) int32 {
	// 先将a视作无符号类型，在转换回来
	return int32(uint32(a) >> (b * 0x1f))
}

func (ins *IOp) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	if ins.isDiv && v2 == 0 {
		// TODO: design runtime exception
		panic("/ by zero.")
	} else {
		frame.PushInt(ins.op(v1, v2))
	}
}

type INeg struct{ base.NoOperandsInstruction }

func (ins *INeg) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	frame.PushInt(-val)
}

// Increment local variable by constant
type IInc struct {
	Index uint
	Const int32
}

func (instr *IInc) FetchOperands(reader *base.BytecodeReader) {
	instr.Index = uint(reader.ReadUint8())
	instr.Const = int32(reader.ReadInt8())
}

func (instr *IInc) Execute(frame *rtda.Frame) {
	val := frame.GetInt(instr.Index)
	val += instr.Const
	frame.SetInt(instr.Index, val)
}
