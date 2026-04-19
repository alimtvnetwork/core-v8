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

// IfFloat64 is a typed convenience wrapper for If[float64].
func IfFloat64(
	isTrue bool,
	trueValue, falseValue float64,
) float64 {
	return If[float64](isTrue, trueValue, falseValue)
}

// IfFuncFloat64 is a typed convenience wrapper for IfFunc[float64].
func IfFuncFloat64(
	isTrue bool,
	trueValueFunc, falseValueFunc func() float64,
) float64 {
	return IfFunc[float64](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncFloat64 is a typed convenience wrapper for IfTrueFunc[float64].
func IfTrueFuncFloat64(
	isTrue bool,
	trueValueFunc func() float64,
) float64 {
	return IfTrueFunc[float64](isTrue, trueValueFunc)
}

// IfSliceFloat64 is a typed convenience wrapper for IfSlice[float64].
func IfSliceFloat64(
	isTrue bool,
	trueValue, falseValue []float64,
) []float64 {
	return IfSlice[float64](isTrue, trueValue, falseValue)
}

// IfPtrFloat64 is a typed convenience wrapper for IfPtr[float64].
func IfPtrFloat64(
	isTrue bool,
	trueValue, falseValue *float64,
) *float64 {
	return IfPtr[float64](isTrue, trueValue, falseValue)
}

// NilDefFloat64 is a typed convenience wrapper for NilDef[float64].
func NilDefFloat64(
	valuePointer *float64,
	defVal float64,
) float64 {
	return NilDef[float64](valuePointer, defVal)
}

// NilDefPtrFloat64 is a typed convenience wrapper for NilDefPtr[float64].
func NilDefPtrFloat64(
	valuePointer *float64,
	defVal float64,
) *float64 {
	return NilDefPtr[float64](valuePointer, defVal)
}

// ValueOrZeroFloat64 is a typed convenience wrapper for ValueOrZero[float64].
func ValueOrZeroFloat64(valuePointer *float64) float64 {
	return ValueOrZero[float64](valuePointer)
}

// PtrOrZeroFloat64 is a typed convenience wrapper for PtrOrZero[float64].
func PtrOrZeroFloat64(valuePointer *float64) *float64 {
	return PtrOrZero[float64](valuePointer)
}

// NilValFloat64 is a typed convenience wrapper for NilVal[float64].
func NilValFloat64(valuePointer *float64, onNil, onNonNil float64) float64 {
	return NilVal[float64](valuePointer, onNil, onNonNil)
}

// NilValPtrFloat64 is a typed convenience wrapper for NilValPtr[float64].
func NilValPtrFloat64(valuePointer *float64, onNil, onNonNil float64) *float64 {
	return NilValPtr[float64](valuePointer, onNil, onNonNil)
}
