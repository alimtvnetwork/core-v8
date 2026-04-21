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
	"fmt"

	"github.com/alimtvnetwork/core-v8/constants"
)

type BaseRange struct {
	RawInput         string
	Separator        string
	IsValid          bool
	HasStart, HasEnd bool
}

func (it *BaseRange) CreateRangeInt(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(
		it.RawInput,
		it.Separator,
		minMax)
}

func (it *BaseRange) IsInvalid() bool {
	return !it.IsValid
}

func (it *BaseRange) BaseRangeClone() *BaseRange {
	return &BaseRange{
		RawInput:  it.RawInput,
		Separator: it.Separator,
		IsValid:   it.IsValid,
		HasStart:  it.HasStart,
		HasEnd:    it.HasEnd,
	}
}

func (it *BaseRange) String(start, end any) string {
	format := constants.SprintValueFormat +
		it.Separator +
		constants.SprintValueFormat

	return fmt.Sprintf(format, start, end)
}
