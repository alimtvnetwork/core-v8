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

// IfFloat32 is a typed convenience wrapper for If[float32].
func IfFloat32(
	isTrue bool,
	trueValue, falseValue float32,
) float32 {
	return If[float32](isTrue, trueValue, falseValue)
}

// IfFuncFloat32 is a typed convenience wrapper for IfFunc[float32].
func IfFuncFloat32(
	isTrue bool,
	trueValueFunc, falseValueFunc func() float32,
) float32 {
	return IfFunc[float32](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncFloat32 is a typed convenience wrapper for IfTrueFunc[float32].
func IfTrueFuncFloat32(
	isTrue bool,
	trueValueFunc func() float32,
) float32 {
	return IfTrueFunc[float32](isTrue, trueValueFunc)
}

// IfSliceFloat32 is a typed convenience wrapper for IfSlice[float32].
func IfSliceFloat32(
	isTrue bool,
	trueValue, falseValue []float32,
) []float32 {
	return IfSlice[float32](isTrue, trueValue, falseValue)
}

// IfPtrFloat32 is a typed convenience wrapper for IfPtr[float32].
func IfPtrFloat32(
	isTrue bool,
	trueValue, falseValue *float32,
) *float32 {
	return IfPtr[float32](isTrue, trueValue, falseValue)
}

// NilDefFloat32 is a typed convenience wrapper for NilDef[float32].
func NilDefFloat32(
	valuePointer *float32,
	defVal float32,
) float32 {
	return NilDef[float32](valuePointer, defVal)
}

// NilDefPtrFloat32 is a typed convenience wrapper for NilDefPtr[float32].
func NilDefPtrFloat32(
	valuePointer *float32,
	defVal float32,
) *float32 {
	return NilDefPtr[float32](valuePointer, defVal)
}

// ValueOrZeroFloat32 is a typed convenience wrapper for ValueOrZero[float32].
func ValueOrZeroFloat32(valuePointer *float32) float32 {
	return ValueOrZero[float32](valuePointer)
}

// PtrOrZeroFloat32 is a typed convenience wrapper for PtrOrZero[float32].
func PtrOrZeroFloat32(valuePointer *float32) *float32 {
	return PtrOrZero[float32](valuePointer)
}

// NilValFloat32 is a typed convenience wrapper for NilVal[float32].
func NilValFloat32(valuePointer *float32, onNil, onNonNil float32) float32 {
	return NilVal[float32](valuePointer, onNil, onNonNil)
}

// NilValPtrFloat32 is a typed convenience wrapper for NilValPtr[float32].
func NilValPtrFloat32(valuePointer *float32, onNil, onNonNil float32) *float32 {
	return NilValPtr[float32](valuePointer, onNil, onNonNil)
}
