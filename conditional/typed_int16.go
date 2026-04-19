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

// IfInt16 is a typed convenience wrapper for If[int16].
func IfInt16(
	isTrue bool,
	trueValue, falseValue int16,
) int16 {
	return If[int16](isTrue, trueValue, falseValue)
}

// IfFuncInt16 is a typed convenience wrapper for IfFunc[int16].
func IfFuncInt16(
	isTrue bool,
	trueValueFunc, falseValueFunc func() int16,
) int16 {
	return IfFunc[int16](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncInt16 is a typed convenience wrapper for IfTrueFunc[int16].
func IfTrueFuncInt16(
	isTrue bool,
	trueValueFunc func() int16,
) int16 {
	return IfTrueFunc[int16](isTrue, trueValueFunc)
}

// IfSliceInt16 is a typed convenience wrapper for IfSlice[int16].
func IfSliceInt16(
	isTrue bool,
	trueValue, falseValue []int16,
) []int16 {
	return IfSlice[int16](isTrue, trueValue, falseValue)
}

// IfPtrInt16 is a typed convenience wrapper for IfPtr[int16].
func IfPtrInt16(
	isTrue bool,
	trueValue, falseValue *int16,
) *int16 {
	return IfPtr[int16](isTrue, trueValue, falseValue)
}

// NilDefInt16 is a typed convenience wrapper for NilDef[int16].
func NilDefInt16(
	valuePointer *int16,
	defVal int16,
) int16 {
	return NilDef[int16](valuePointer, defVal)
}

// NilDefPtrInt16 is a typed convenience wrapper for NilDefPtr[int16].
func NilDefPtrInt16(
	valuePointer *int16,
	defVal int16,
) *int16 {
	return NilDefPtr[int16](valuePointer, defVal)
}

// ValueOrZeroInt16 is a typed convenience wrapper for ValueOrZero[int16].
func ValueOrZeroInt16(valuePointer *int16) int16 {
	return ValueOrZero[int16](valuePointer)
}

// PtrOrZeroInt16 is a typed convenience wrapper for PtrOrZero[int16].
func PtrOrZeroInt16(valuePointer *int16) *int16 {
	return PtrOrZero[int16](valuePointer)
}

// NilValInt16 is a typed convenience wrapper for NilVal[int16].
func NilValInt16(valuePointer *int16, onNil, onNonNil int16) int16 {
	return NilVal[int16](valuePointer, onNil, onNonNil)
}

// NilValPtrInt16 is a typed convenience wrapper for NilValPtr[int16].
func NilValPtrInt16(valuePointer *int16, onNil, onNonNil int16) *int16 {
	return NilValPtr[int16](valuePointer, onNil, onNonNil)
}
