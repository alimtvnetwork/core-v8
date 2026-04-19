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

// IfUint8 is a typed convenience wrapper for If[uint8].
func IfUint8(
	isTrue bool,
	trueValue, falseValue uint8,
) uint8 {
	return If[uint8](isTrue, trueValue, falseValue)
}

// IfFuncUint8 is a typed convenience wrapper for IfFunc[uint8].
func IfFuncUint8(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint8,
) uint8 {
	return IfFunc[uint8](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint8 is a typed convenience wrapper for IfTrueFunc[uint8].
func IfTrueFuncUint8(
	isTrue bool,
	trueValueFunc func() uint8,
) uint8 {
	return IfTrueFunc[uint8](isTrue, trueValueFunc)
}

// IfSliceUint8 is a typed convenience wrapper for IfSlice[uint8].
func IfSliceUint8(
	isTrue bool,
	trueValue, falseValue []uint8,
) []uint8 {
	return IfSlice[uint8](isTrue, trueValue, falseValue)
}

// IfPtrUint8 is a typed convenience wrapper for IfPtr[uint8].
func IfPtrUint8(
	isTrue bool,
	trueValue, falseValue *uint8,
) *uint8 {
	return IfPtr[uint8](isTrue, trueValue, falseValue)
}

// NilDefUint8 is a typed convenience wrapper for NilDef[uint8].
func NilDefUint8(
	valuePointer *uint8,
	defVal uint8,
) uint8 {
	return NilDef[uint8](valuePointer, defVal)
}

// NilDefPtrUint8 is a typed convenience wrapper for NilDefPtr[uint8].
func NilDefPtrUint8(
	valuePointer *uint8,
	defVal uint8,
) *uint8 {
	return NilDefPtr[uint8](valuePointer, defVal)
}

// ValueOrZeroUint8 is a typed convenience wrapper for ValueOrZero[uint8].
func ValueOrZeroUint8(valuePointer *uint8) uint8 {
	return ValueOrZero[uint8](valuePointer)
}

// PtrOrZeroUint8 is a typed convenience wrapper for PtrOrZero[uint8].
func PtrOrZeroUint8(valuePointer *uint8) *uint8 {
	return PtrOrZero[uint8](valuePointer)
}

// NilValUint8 is a typed convenience wrapper for NilVal[uint8].
func NilValUint8(valuePointer *uint8, onNil, onNonNil uint8) uint8 {
	return NilVal[uint8](valuePointer, onNil, onNonNil)
}

// NilValPtrUint8 is a typed convenience wrapper for NilValPtr[uint8].
func NilValPtrUint8(valuePointer *uint8, onNil, onNonNil uint8) *uint8 {
	return NilValPtr[uint8](valuePointer, onNil, onNonNil)
}
