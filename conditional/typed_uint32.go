// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package conditional

// IfUint32 is a typed convenience wrapper for If[uint32].
func IfUint32(
	isTrue bool,
	trueValue, falseValue uint32,
) uint32 {
	return If[uint32](isTrue, trueValue, falseValue)
}

// IfFuncUint32 is a typed convenience wrapper for IfFunc[uint32].
func IfFuncUint32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint32,
) uint32 {
	return IfFunc[uint32](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint32 is a typed convenience wrapper for IfTrueFunc[uint32].
func IfTrueFuncUint32(
	isTrue bool,
	trueValueFunc func() uint32,
) uint32 {
	return IfTrueFunc[uint32](isTrue, trueValueFunc)
}

// IfSliceUint32 is a typed convenience wrapper for IfSlice[uint32].
func IfSliceUint32(
	isTrue bool,
	trueValue, falseValue []uint32,
) []uint32 {
	return IfSlice[uint32](isTrue, trueValue, falseValue)
}

// IfPtrUint32 is a typed convenience wrapper for IfPtr[uint32].
func IfPtrUint32(
	isTrue bool,
	trueValue, falseValue *uint32,
) *uint32 {
	return IfPtr[uint32](isTrue, trueValue, falseValue)
}

// NilDefUint32 is a typed convenience wrapper for NilDef[uint32].
func NilDefUint32(
	valuePointer *uint32,
	defVal uint32,
) uint32 {
	return NilDef[uint32](valuePointer, defVal)
}

// NilDefPtrUint32 is a typed convenience wrapper for NilDefPtr[uint32].
func NilDefPtrUint32(
	valuePointer *uint32,
	defVal uint32,
) *uint32 {
	return NilDefPtr[uint32](valuePointer, defVal)
}

// ValueOrZeroUint32 is a typed convenience wrapper for ValueOrZero[uint32].
func ValueOrZeroUint32(valuePointer *uint32) uint32 {
	return ValueOrZero[uint32](valuePointer)
}

// PtrOrZeroUint32 is a typed convenience wrapper for PtrOrZero[uint32].
func PtrOrZeroUint32(valuePointer *uint32) *uint32 {
	return PtrOrZero[uint32](valuePointer)
}

// NilValUint32 is a typed convenience wrapper for NilVal[uint32].
func NilValUint32(valuePointer *uint32, onNil, onNonNil uint32) uint32 {
	return NilVal[uint32](valuePointer, onNil, onNonNil)
}

// NilValPtrUint32 is a typed convenience wrapper for NilValPtr[uint32].
func NilValPtrUint32(valuePointer *uint32, onNil, onNonNil uint32) *uint32 {
	return NilValPtr[uint32](valuePointer, onNil, onNonNil)
}
