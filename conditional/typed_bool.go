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

// IfBool is a typed convenience wrapper for If[bool].
func IfBool(
	isTrue bool,
	trueValue, falseValue bool,
) bool {
	return If[bool](isTrue, trueValue, falseValue)
}

// IfFuncBool is a typed convenience wrapper for IfFunc[bool].
func IfFuncBool(
	isTrue bool,
	trueValueFunc, falseValueFunc func() bool,
) bool {
	return IfFunc[bool](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncBool is a typed convenience wrapper for IfTrueFunc[bool].
func IfTrueFuncBool(
	isTrue bool,
	trueValueFunc func() bool,
) bool {
	return IfTrueFunc[bool](isTrue, trueValueFunc)
}

// IfSliceBool is a typed convenience wrapper for IfSlice[bool].
func IfSliceBool(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}

// IfPtrBool is a typed convenience wrapper for IfPtr[bool].
func IfPtrBool(
	isTrue bool,
	trueValue, falseValue *bool,
) *bool {
	return IfPtr[bool](isTrue, trueValue, falseValue)
}

// NilDefBool is a typed convenience wrapper for NilDef[bool].
func NilDefBool(
	valuePointer *bool,
	defVal bool,
) bool {
	return NilDef[bool](valuePointer, defVal)
}

// NilDefPtrBool is a typed convenience wrapper for NilDefPtr[bool].
func NilDefPtrBool(
	valuePointer *bool,
	defVal bool,
) *bool {
	return NilDefPtr[bool](valuePointer, defVal)
}

// ValueOrZeroBool is a typed convenience wrapper for ValueOrZero[bool].
func ValueOrZeroBool(valuePointer *bool) bool {
	return ValueOrZero[bool](valuePointer)
}

// PtrOrZeroBool is a typed convenience wrapper for PtrOrZero[bool].
func PtrOrZeroBool(valuePointer *bool) *bool {
	return PtrOrZero[bool](valuePointer)
}

// NilValBool is a typed convenience wrapper for NilVal[bool].
func NilValBool(valuePointer *bool, onNil, onNonNil bool) bool {
	return NilVal[bool](valuePointer, onNil, onNonNil)
}

// NilValPtrBool is a typed convenience wrapper for NilValPtr[bool].
func NilValPtrBool(valuePointer *bool, onNil, onNonNil bool) *bool {
	return NilValPtr[bool](valuePointer, onNil, onNonNil)
}
