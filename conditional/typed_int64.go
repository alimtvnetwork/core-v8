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

// IfInt64 is a typed convenience wrapper for If[int64].
func IfInt64(
	isTrue bool,
	trueValue, falseValue int64,
) int64 {
	return If[int64](isTrue, trueValue, falseValue)
}

// IfFuncInt64 is a typed convenience wrapper for IfFunc[int64].
func IfFuncInt64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int64,
) int64 {
	return IfFunc[int64](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt64 is a typed convenience wrapper for IfTrueFunc[int64].
func IfTrueFuncInt64(
	isTrue bool,
	trueValueFunc func() int64,
) int64 {
	return IfTrueFunc[int64](isTrue, trueValueFunc)
}

// IfSliceInt64 is a typed convenience wrapper for IfSlice[int64].
func IfSliceInt64(
	isTrue bool,
	trueValue, falseValue []int64,
) []int64 {
	return IfSlice[int64](isTrue, trueValue, falseValue)
}

// IfPtrInt64 is a typed convenience wrapper for IfPtr[int64].
func IfPtrInt64(
	isTrue bool,
	trueValue, falseValue *int64,
) *int64 {
	return IfPtr[int64](isTrue, trueValue, falseValue)
}

// NilDefInt64 is a typed convenience wrapper for NilDef[int64].
func NilDefInt64(
	valuePointer *int64,
	defVal int64,
) int64 {
	return NilDef[int64](valuePointer, defVal)
}

// NilDefPtrInt64 is a typed convenience wrapper for NilDefPtr[int64].
func NilDefPtrInt64(
	valuePointer *int64,
	defVal int64,
) *int64 {
	return NilDefPtr[int64](valuePointer, defVal)
}

// ValueOrZeroInt64 is a typed convenience wrapper for ValueOrZero[int64].
func ValueOrZeroInt64(valuePointer *int64) int64 {
	return ValueOrZero[int64](valuePointer)
}

// PtrOrZeroInt64 is a typed convenience wrapper for PtrOrZero[int64].
func PtrOrZeroInt64(valuePointer *int64) *int64 {
	return PtrOrZero[int64](valuePointer)
}

// NilValInt64 is a typed convenience wrapper for NilVal[int64].
func NilValInt64(valuePointer *int64, onNil, onNonNil int64) int64 {
	return NilVal[int64](valuePointer, onNil, onNonNil)
}

// NilValPtrInt64 is a typed convenience wrapper for NilValPtr[int64].
func NilValPtrInt64(valuePointer *int64, onNil, onNonNil int64) *int64 {
	return NilValPtr[int64](valuePointer, onNil, onNonNil)
}
