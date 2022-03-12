package base

import "jvmgo/ch06/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (ins *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// do nothing
}

type BranchInstruction struct {
	Offset int
}

func (branch *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	branch.Offset = int(reader.ReadInt16())
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.aload
// E.g aload index;
// The index is an unsigned byte that must be an index into the local variable array of the current frame
type Index8Instruction struct {
	Index uint
}

func (idx8 *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	idx8.Index = uint(reader.ReadUint8())
}

// E.g ldc_w index;
// The unsigned indexbyte1 and indexbyte2 are assembled
// into an unsigned 16-bit index into the run-time constant pool of the current class
type Index16Instruction struct {
	Index uint
}

func (index16 *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	index16.Index = uint(reader.ReadUint16())
}
