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

// IfString is a typed convenience wrapper for If[string].
func IfString(
	isTrue bool,
	trueValue, falseValue string,
) string {
	return If[string](isTrue, trueValue, falseValue)
}

// IfFuncString is a typed convenience wrapper for IfFunc[string].
func IfFuncString(
	isTrue bool,
	trueValueFunc, falseValueFunc func() string,
) string {
	return IfFunc[string](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncString is a typed convenience wrapper for IfTrueFunc[string].
func IfTrueFuncString(
	isTrue bool,
	trueValueFunc func() string,
) string {
	return IfTrueFunc[string](isTrue, trueValueFunc)
}

// IfSliceString is a typed convenience wrapper for IfSlice[string].
func IfSliceString(
	isTrue bool,
	trueValue, falseValue []string,
) []string {
	return IfSlice[string](isTrue, trueValue, falseValue)
}

// IfPtrString is a typed convenience wrapper for IfPtr[string].
func IfPtrString(
	isTrue bool,
	trueValue, falseValue *string,
) *string {
	return IfPtr[string](isTrue, trueValue, falseValue)
}

// NilDefString is a typed convenience wrapper for NilDef[string].
func NilDefString(
	valuePointer *string,
	defVal string,
) string {
	return NilDef[string](valuePointer, defVal)
}

// NilDefPtrString is a typed convenience wrapper for NilDefPtr[string].
func NilDefPtrString(
	valuePointer *string,
	defVal string,
) *string {
	return NilDefPtr[string](valuePointer, defVal)
}

// ValueOrZeroString is a typed convenience wrapper for ValueOrZero[string].
func ValueOrZeroString(valuePointer *string) string {
	return ValueOrZero[string](valuePointer)
}

// PtrOrZeroString is a typed convenience wrapper for PtrOrZero[string].
func PtrOrZeroString(valuePointer *string) *string {
	return PtrOrZero[string](valuePointer)
}

// NilValString is a typed convenience wrapper for NilVal[string].
func NilValString(valuePointer *string, onNil, onNonNil string) string {
	return NilVal[string](valuePointer, onNil, onNonNil)
}

// NilValPtrString is a typed convenience wrapper for NilValPtr[string].
func NilValPtrString(valuePointer *string, onNil, onNonNil string) *string {
	return NilValPtr[string](valuePointer, onNil, onNonNil)
}
