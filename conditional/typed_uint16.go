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

// IfUint16 is a typed convenience wrapper for If[uint16].
func IfUint16(
	isTrue bool,
	trueValue, falseValue uint16,
) uint16 {
	return If[uint16](isTrue, trueValue, falseValue)
}

// IfFuncUint16 is a typed convenience wrapper for IfFunc[uint16].
func IfFuncUint16(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint16,
) uint16 {
	return IfFunc[uint16](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint16 is a typed convenience wrapper for IfTrueFunc[uint16].
func IfTrueFuncUint16(
	isTrue bool,
	trueValueFunc func() uint16,
) uint16 {
	return IfTrueFunc[uint16](isTrue, trueValueFunc)
}

// IfSliceUint16 is a typed convenience wrapper for IfSlice[uint16].
func IfSliceUint16(
	isTrue bool,
	trueValue, falseValue []uint16,
) []uint16 {
	return IfSlice[uint16](isTrue, trueValue, falseValue)
}

// IfPtrUint16 is a typed convenience wrapper for IfPtr[uint16].
func IfPtrUint16(
	isTrue bool,
	trueValue, falseValue *uint16,
) *uint16 {
	return IfPtr[uint16](isTrue, trueValue, falseValue)
}

// NilDefUint16 is a typed convenience wrapper for NilDef[uint16].
func NilDefUint16(
	valuePointer *uint16,
	defVal uint16,
) uint16 {
	return NilDef[uint16](valuePointer, defVal)
}

// NilDefPtrUint16 is a typed convenience wrapper for NilDefPtr[uint16].
func NilDefPtrUint16(
	valuePointer *uint16,
	defVal uint16,
) *uint16 {
	return NilDefPtr[uint16](valuePointer, defVal)
}

// ValueOrZeroUint16 is a typed convenience wrapper for ValueOrZero[uint16].
func ValueOrZeroUint16(valuePointer *uint16) uint16 {
	return ValueOrZero[uint16](valuePointer)
}

// PtrOrZeroUint16 is a typed convenience wrapper for PtrOrZero[uint16].
func PtrOrZeroUint16(valuePointer *uint16) *uint16 {
	return PtrOrZero[uint16](valuePointer)
}

// NilValUint16 is a typed convenience wrapper for NilVal[uint16].
func NilValUint16(valuePointer *uint16, onNil, onNonNil uint16) uint16 {
	return NilVal[uint16](valuePointer, onNil, onNonNil)
}

// NilValPtrUint16 is a typed convenience wrapper for NilValPtr[uint16].
func NilValPtrUint16(valuePointer *uint16, onNil, onNonNil uint16) *uint16 {
	return NilValPtr[uint16](valuePointer, onNil, onNonNil)
}
