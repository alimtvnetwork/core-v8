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

package corerange

import (
	"github.com/alimtvnetwork/core-v8/constants"
)

type RangeByte struct {
	*BaseRange
	Start, End byte
}

func NewRangeByteMinMax(
	rawString, separator string,
	min, max byte,
) *RangeByte {
	minMax := MinMaxInt{
		Min: int(min),
		Max: int(max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMax).
		CreateRangeByte()
}

func NewRangeByte(
	rawString, separator string,
	minMax *MinMaxByte,
) *RangeByte {
	if minMax == nil {
		minMaxInt := MinMaxInt{
			Min: constants.Zero,
			Max: constants.MaxUnit8AsInt,
		}

		return NewRangeInt(
			rawString,
			separator,
			&minMaxInt).
			CreateRangeByte()
	}

	minMaxInt := MinMaxInt{
		Min: int(minMax.Min),
		Max: int(minMax.Max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMaxInt).
		CreateRangeByte()
}

// Difference
//
// Checks comparison wise which one is bigger than does the diff.
func (it *RangeByte) Difference() byte {
	if it.End > it.Start {
		return it.End - it.Start
	}

	return it.Start - it.End
}

func (it *RangeByte) DifferenceAbsolute() byte {
	return it.Difference()
}

// RangeLength (5 - 3 = 2) + 1
func (it *RangeByte) RangeLength() byte {
	return it.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (it *RangeByte) RangesInt() []byte {
	return it.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *RangeByte) Ranges() []byte {
	if it.IsInvalid() {
		return []byte{}
	}

	length := int(it.DifferenceAbsolute()) + 1
	start := it.Start
	slice := make([]byte, length)

	for i := 0; i < length; i++ {
		slice[i] = start + byte(i)
	}

	return slice
}

// IsWithinRange it.Start <= value && value <= it.End
func (it *RangeByte) IsWithinRange(value byte) bool {
	return it.Start <= value && value <= it.End
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (it *RangeByte) IsValidPlusWithinRange(value byte) bool {
	return it.IsValid && it.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (it *RangeByte) IsInvalidValue(value byte) bool {
	return !it.IsValid || !it.IsWithinRange(value)
}

func (it *RangeByte) String() string {
	return it.BaseRange.String(it.Start, it.End)
}
