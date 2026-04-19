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

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/osconsts"
)

type integerOutOfRange struct{}

func (it integerOutOfRange) ToByte(value int) bool {
	return !(value >= 0 && value <= 255)
}

func (it integerOutOfRange) ToUnsignedInt16(value int) bool {
	return !(value >= 0 && value <= math.MaxUint16)
}

func (it integerOutOfRange) ToUnsignedInt32(value int) bool {
	if osconsts.IsX32Architecture {
		return !(value >= 0 && value <= math.MaxInt32)
	}

	return !(value >= 0 && value <= math.MaxUint32)
}

func (it integerOutOfRange) ToUnsignedInt64(value int) bool {
	return !(value >= 0)
}

func (it integerOutOfRange) ToInt8(value int) bool {
	return !(value >= math.MinInt8 && value <= math.MaxInt8)
}

func (it integerOutOfRange) ToInt16(value int) bool {
	return !(value >= math.MinInt16 && value <= math.MaxInt16)
}

func (it integerOutOfRange) ToInt32(value int) bool {
	return !(value >= math.MinInt32 && value <= math.MaxInt32)
}

func (it integerOutOfRange) ToInt(value int) bool {
	return !(value >= constants.MinInt && value <= constants.MaxInt)
}
