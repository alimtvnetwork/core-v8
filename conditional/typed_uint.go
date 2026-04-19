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

// IfUint is a typed convenience wrapper for If[uint].
func IfUint(
	isTrue bool,
	trueValue, falseValue uint,
) uint {
	return If[uint](isTrue, trueValue, falseValue)
}

// IfFuncUint is a typed convenience wrapper for IfFunc[uint].
func IfFuncUint(
	isTrue bool,
	trueValueFunc, falseValueFunc func() uint,
) uint {
	return IfFunc[uint](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncUint is a typed convenience wrapper for IfTrueFunc[uint].
func IfTrueFuncUint(
	isTrue bool,
	trueValueFunc func() uint,
) uint {
	return IfTrueFunc[uint](isTrue, trueValueFunc)
}

// IfSliceUint is a typed convenience wrapper for IfSlice[uint].
func IfSliceUint(
	isTrue bool,
	trueValue, falseValue []uint,
) []uint {
	return IfSlice[uint](isTrue, trueValue, falseValue)
}

// IfPtrUint is a typed convenience wrapper for IfPtr[uint].
func IfPtrUint(
	isTrue bool,
	trueValue, falseValue *uint,
) *uint {
	return IfPtr[uint](isTrue, trueValue, falseValue)
}

// NilDefUint is a typed convenience wrapper for NilDef[uint].
func NilDefUint(
	valuePointer *uint,
	defVal uint,
) uint {
	return NilDef[uint](valuePointer, defVal)
}

// NilDefPtrUint is a typed convenience wrapper for NilDefPtr[uint].
func NilDefPtrUint(
	valuePointer *uint,
	defVal uint,
) *uint {
	return NilDefPtr[uint](valuePointer, defVal)
}

// ValueOrZeroUint is a typed convenience wrapper for ValueOrZero[uint].
func ValueOrZeroUint(valuePointer *uint) uint {
	return ValueOrZero[uint](valuePointer)
}

// PtrOrZeroUint is a typed convenience wrapper for PtrOrZero[uint].
func PtrOrZeroUint(valuePointer *uint) *uint {
	return PtrOrZero[uint](valuePointer)
}

// NilValUint is a typed convenience wrapper for NilVal[uint].
func NilValUint(valuePointer *uint, onNil, onNonNil uint) uint {
	return NilVal[uint](valuePointer, onNil, onNonNil)
}

// NilValPtrUint is a typed convenience wrapper for NilValPtr[uint].
func NilValPtrUint(valuePointer *uint, onNil, onNonNil uint) *uint {
	return NilValPtr[uint](valuePointer, onNil, onNonNil)
}
