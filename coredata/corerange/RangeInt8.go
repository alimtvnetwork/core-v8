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
	"math"
)

type RangeInt8 struct {
	*BaseRange
	Start, End int8
}

func NewRangeInt8MinMax(
	rawString, separator string,
	min, max int8,
) *RangeInt8 {
	minMax := MinMaxInt{
		Min: int(min),
		Max: int(max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMax).
		CreateRangeInt8()
}

func NewRangeInt8(
	rawString, separator string,
	minMaxInt8 *MinMaxInt8,
) *RangeInt8 {
	if minMaxInt8 == nil {
		minMax := &MinMaxInt{
			Min: math.MinInt8,
			Max: math.MaxInt8,
		}

		rangeInt := NewRangeInt(
			rawString,
			separator,
			minMax)

		return rangeInt.CreateRangeInt8()
	}

	minMax := minMaxInt8.CreateMinMaxInt()

	rangeInt := NewRangeInt(
		rawString,
		separator,
		minMax)

	return rangeInt.CreateRangeInt8()
}

func (it *RangeInt8) Difference() int8 {
	return it.End - it.Start
}

func (it *RangeInt8) DifferenceAbsolute() int8 {
	diff := it.Difference()

	if diff < 0 {
		return diff * -1
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (it *RangeInt8) RangeLength() int8 {
	return it.DifferenceAbsolute() + 1
}

// RangesInt8 returns empty ints if IsInvalid
// return range int values
func (it *RangeInt8) RangesInt8() []int8 {
	return it.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *RangeInt8) Ranges() []int8 {
	if it.IsInvalid() {
		return []int8{}
	}

	length := it.RangeLength()
	start := it.Start
	slice := make([]int8, length)

	var i int8

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *RangeInt8) String() string {
	return it.BaseRange.String(it.Start, it.End)
}

// IsWithinRange it.End >= value && value >= it.Start
func (it *RangeInt8) IsWithinRange(value int8) bool {
	return it.End >= value && value >= it.Start
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (it *RangeInt8) IsValidPlusWithinRange(value int8) bool {
	return it.IsValid && it.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (it *RangeInt8) IsInvalidValue(value int8) bool {
	return !it.IsValid || !it.IsWithinRange(value)
}
