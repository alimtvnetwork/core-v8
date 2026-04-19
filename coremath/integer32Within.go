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

package coremath

import (
	"math"
)

type integer32Within struct{}

func (it integer32Within) ToByte(value int32) bool {
	return value >= 0 && value <= 255
}

func (it integer32Within) ToUnsignedInt16(value int32) bool {
	return value >= 0 && value <= math.MaxUint16
}

func (it integer32Within) ToUnsignedInt32(value int32) bool {
	return value >= 0
}

func (it integer32Within) ToUnsignedInt64(value int32) bool {
	return value >= 0
}

func (it integer32Within) ToInt8(value int32) bool {
	return value >= math.MinInt8 && value <= math.MaxInt8
}

func (it integer32Within) ToInt16(value int32) bool {
	return value >= math.MinInt16 && value <= math.MaxInt16
}

func (it integer32Within) ToInt(value int) bool {
	// it would be different based on architecture.
	return value >= math.MinInt32 && value <= math.MaxInt32
}
