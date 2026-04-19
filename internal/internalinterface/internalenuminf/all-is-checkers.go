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

package internalenuminf

type IsValidChecker interface {
	// IsValid similar or alias for IsSuccessChecker
	IsValid() bool
}

type IsInvalidChecker interface {
	IsInvalid() bool
}

type IsValidInvalidChecker interface {
	IsValidChecker
	IsInvalidChecker
}

type IsNameEqualer interface {
	IsNameEqual(name string) bool
}

// IsAnyNameOfChecker
//
//	Returns true if any of the name matches.
type IsAnyNameOfChecker interface {
	// IsAnyNamesOf
	//
	//  Returns true if any of the name matches.
	IsAnyNamesOf(names ...string) bool
}

type IsAnyValueByteEqualer interface {
	IsAnyValuesEqual(anyByteValues ...byte) bool
}

type IsAnyValueIntegerEqualer interface {
	IsAnyValuesEqual(anyValues ...int) bool
}

type IsAnyValueInteger8Equaler interface {
	IsAnyValuesEqual(anyValues ...int8) bool
}

type IsAnyValueInteger16Equaler interface {
	IsAnyValuesEqual(anyValues ...int16) bool
}

type IsAnyValueInteger32Equaler interface {
	IsAnyValuesEqual(anyValues ...int32) bool
}

type IsByteValueEqualer interface {
	IsByteValueEqual(value byte) bool
}

type IsValueIntegerEqualer interface {
	IsValueEqual(value int) bool
}

type IsValueInteger8Equaler interface {
	IsValueEqual(value int8) bool
}

type IsValueInteger16Equaler interface {
	IsValueEqual(value int16) bool
}

type IsValueInteger32Equaler interface {
	IsValueEqual(value int32) bool
}

type RangeValidateChecker interface {
	// RangesInvalidMessage get invalid message
	RangesInvalidMessage() string
	// RangesInvalidErr get invalid message error
	RangesInvalidErr() error
	// IsValidRange Is with in the range as expected.
	IsValidRange() bool
	// IsInvalidRange Is out of the ranges expected.
	IsInvalidRange() bool
}

type IsEnumEqualer interface {
	IsEnumEqual(enum BasicEnumer) bool
}

type IsAnyEnumsEqualer interface {
	IsAnyEnumsEqual(enums ...BasicEnumer) bool
}
