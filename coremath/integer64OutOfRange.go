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
)

type integer64OutOfRange struct{}

func (it integer64OutOfRange) Byte(value int64) bool {
	return !(value >= 0 && value <= 255)
}

func (it integer64OutOfRange) UnsignedInt16(value int64) bool {
	return !(value >= 0 && value <= int64(math.MaxUint16))
}

func (it integer64OutOfRange) UnsignedInt32(value int64) bool {
	return !(value >= 0 && value <= int64(math.MaxUint32))
}

func (it integer64OutOfRange) UnsignedInt64(value int64) bool {
	return !(value >= 0)
}

func (it integer64OutOfRange) Int8(value int64) bool {
	return !(value >= int64(math.MinInt8) && value <= int64(math.MaxInt8))
}

func (it integer64OutOfRange) Int16(value int64) bool {
	return !(value >= int64(math.MinInt16) && value <= int64(math.MaxInt16))
}

func (it integer64OutOfRange) Int32(value int64) bool {
	return !(value >= int64(math.MinInt32) && value <= int64(math.MaxInt32))
}

func (it integer64OutOfRange) Int(value int64) bool {
	return !(value >= int64(constants.MinInt) && value <= int64(constants.MaxInt))
}
