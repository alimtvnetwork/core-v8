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

// IfUint64 is a typed convenience wrapper for If[uint64].
func IfUint64(
	isTrue bool,
	trueValue, falseValue uint64,
) uint64 {
	return If[uint64](isTrue, trueValue, falseValue)
}

// IfFuncUint64 is a typed convenience wrapper for IfFunc[uint64].
func IfFuncUint64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint64,
) uint64 {
	return IfFunc[uint64](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint64 is a typed convenience wrapper for IfTrueFunc[uint64].
func IfTrueFuncUint64(
	isTrue bool,
	trueValueFunc func() uint64,
) uint64 {
	return IfTrueFunc[uint64](isTrue, trueValueFunc)
}

// IfSliceUint64 is a typed convenience wrapper for IfSlice[uint64].
func IfSliceUint64(
	isTrue bool,
	trueValue, falseValue []uint64,
) []uint64 {
	return IfSlice[uint64](isTrue, trueValue, falseValue)
}

// IfPtrUint64 is a typed convenience wrapper for IfPtr[uint64].
func IfPtrUint64(
	isTrue bool,
	trueValue, falseValue *uint64,
) *uint64 {
	return IfPtr[uint64](isTrue, trueValue, falseValue)
}

// NilDefUint64 is a typed convenience wrapper for NilDef[uint64].
func NilDefUint64(
	valuePointer *uint64,
	defVal uint64,
) uint64 {
	return NilDef[uint64](valuePointer, defVal)
}

// NilDefPtrUint64 is a typed convenience wrapper for NilDefPtr[uint64].
func NilDefPtrUint64(
	valuePointer *uint64,
	defVal uint64,
) *uint64 {
	return NilDefPtr[uint64](valuePointer, defVal)
}

// ValueOrZeroUint64 is a typed convenience wrapper for ValueOrZero[uint64].
func ValueOrZeroUint64(valuePointer *uint64) uint64 {
	return ValueOrZero[uint64](valuePointer)
}

// PtrOrZeroUint64 is a typed convenience wrapper for PtrOrZero[uint64].
func PtrOrZeroUint64(valuePointer *uint64) *uint64 {
	return PtrOrZero[uint64](valuePointer)
}

// NilValUint64 is a typed convenience wrapper for NilVal[uint64].
func NilValUint64(valuePointer *uint64, onNil, onNonNil uint64) uint64 {
	return NilVal[uint64](valuePointer, onNil, onNonNil)
}

// NilValPtrUint64 is a typed convenience wrapper for NilValPtr[uint64].
func NilValPtrUint64(valuePointer *uint64, onNil, onNonNil uint64) *uint64 {
	return NilValPtr[uint64](valuePointer, onNil, onNonNil)
}
