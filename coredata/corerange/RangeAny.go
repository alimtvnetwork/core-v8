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

import "github.com/alimtvnetwork/core-v8/internal/strutilinternal"

type RangeAny struct {
	*BaseRange
	RawInput   any
	Start, End any
}

func (r *RangeAny) RawInputString() string {
	return strutilinternal.AnyToString(r.RawInput)
}

func (r *RangeAny) StartString() string {
	return strutilinternal.AnyToString(r.Start)
}

func (r *RangeAny) EndString() string {
	return strutilinternal.AnyToString(r.End)
}

func (r *RangeAny) CreateRangeInt() *RangeInt {
	return NewRangeInt(
		r.RawInputString(),
		r.Separator,
		nil)
}

func (r *RangeAny) CreateRangeIntMinMax(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(
		r.RawInputString(),
		r.Separator,
		minMax)
}

func (r *RangeAny) CreateRangeString() *RangeString {
	return &RangeString{
		StartEndString: r.CreateStartEndString(),
	}
}

func (r *RangeAny) CreateStartEndString() *StartEndString {
	return &StartEndString{
		BaseRange: r.BaseRangeClone(),
		Start:     r.StartString(),
		End:       r.EndString(),
	}
}

func (r *RangeAny) String() string {
	return r.BaseRange.String(r.Start, r.End)
}
