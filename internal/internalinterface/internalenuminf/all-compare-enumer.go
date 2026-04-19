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

type CompareByteEnumer interface {
	IsEqual(comparingValue byte) bool
	IsLess(comparingValue byte) bool
	IsLessEqual(comparingValue byte) bool
	IsGreater(comparingValue byte) bool
	IsGreaterEqual(comparingValue byte) bool
	IsNotEqual(comparingValue byte) bool
}

type CompareInt8Enumer interface {
	IsEqual(comparingValue int8) bool
	IsLess(comparingValue int8) bool
	IsLessEqual(comparingValue int8) bool
	IsGreater(comparingValue int8) bool
	IsGreaterEqual(comparingValue int8) bool
	IsNotEqual(comparingValue int8) bool
}

type CompareInt16Enumer interface {
	IsEqual(comparingValue int16) bool
	IsLess(comparingValue int16) bool
	IsLessEqual(comparingValue int16) bool
	IsGreater(comparingValue int16) bool
	IsGreaterEqual(comparingValue int16) bool
	IsNotEqual(comparingValue int16) bool
}

type CompareInt32Enumer interface {
	IsEqual(comparingValue int32) bool
	IsLess(comparingValue int32) bool
	IsLessEqual(comparingValue int32) bool
	IsGreater(comparingValue int32) bool
	IsGreaterEqual(comparingValue int32) bool
	IsNotEqual(comparingValue int32) bool
}

type CompareIntegerEnumer interface {
	IsEqual(comparingValue int) bool
	IsLess(comparingValue int) bool
	IsLessEqual(comparingValue int) bool
	IsGreater(comparingValue int) bool
	IsGreaterEqual(comparingValue int) bool
	IsNotEqual(comparingValue int) bool
}

type CompareBasicEnumer interface {
	IsEqual(enum BasicEnumer) bool
	IsLess(enum BasicEnumer) bool
	IsLessEqual(enum BasicEnumer) bool
	IsGreater(enum BasicEnumer) bool
	IsGreaterEqual(enum BasicEnumer) bool
	IsNotEqual(enum BasicEnumer) bool
}
