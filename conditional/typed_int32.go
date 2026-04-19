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

// IfInt32 is a typed convenience wrapper for If[int32].
func IfInt32(
	isTrue bool,
	trueValue, falseValue int32,
) int32 {
	return If[int32](isTrue, trueValue, falseValue)
}

// IfFuncInt32 is a typed convenience wrapper for IfFunc[int32].
func IfFuncInt32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int32,
) int32 {
	return IfFunc[int32](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt32 is a typed convenience wrapper for IfTrueFunc[int32].
func IfTrueFuncInt32(
	isTrue bool,
	trueValueFunc func() int32,
) int32 {
	return IfTrueFunc[int32](isTrue, trueValueFunc)
}

// IfSliceInt32 is a typed convenience wrapper for IfSlice[int32].
func IfSliceInt32(
	isTrue bool,
	trueValue, falseValue []int32,
) []int32 {
	return IfSlice[int32](isTrue, trueValue, falseValue)
}

// IfPtrInt32 is a typed convenience wrapper for IfPtr[int32].
func IfPtrInt32(
	isTrue bool,
	trueValue, falseValue *int32,
) *int32 {
	return IfPtr[int32](isTrue, trueValue, falseValue)
}

// NilDefInt32 is a typed convenience wrapper for NilDef[int32].
func NilDefInt32(
	valuePointer *int32,
	defVal int32,
) int32 {
	return NilDef[int32](valuePointer, defVal)
}

// NilDefPtrInt32 is a typed convenience wrapper for NilDefPtr[int32].
func NilDefPtrInt32(
	valuePointer *int32,
	defVal int32,
) *int32 {
	return NilDefPtr[int32](valuePointer, defVal)
}

// ValueOrZeroInt32 is a typed convenience wrapper for ValueOrZero[int32].
func ValueOrZeroInt32(valuePointer *int32) int32 {
	return ValueOrZero[int32](valuePointer)
}

// PtrOrZeroInt32 is a typed convenience wrapper for PtrOrZero[int32].
func PtrOrZeroInt32(valuePointer *int32) *int32 {
	return PtrOrZero[int32](valuePointer)
}

// NilValInt32 is a typed convenience wrapper for NilVal[int32].
func NilValInt32(valuePointer *int32, onNil, onNonNil int32) int32 {
	return NilVal[int32](valuePointer, onNil, onNonNil)
}

// NilValPtrInt32 is a typed convenience wrapper for NilValPtr[int32].
func NilValPtrInt32(valuePointer *int32, onNil, onNonNil int32) *int32 {
	return NilValPtr[int32](valuePointer, onNil, onNonNil)
}
