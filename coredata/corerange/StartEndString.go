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
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coreindexes"
)

type StartEndString struct {
	*BaseRange
	Start, End string
}

func NewStartEndString(raw, sep string) *StartEndString {
	ranges := strings.Split(raw, sep)
	length := len(ranges)
	hasStart := length >= 1
	hasEnd := length >= 2
	isValid := length == 2 &&
		hasStart &&
		hasEnd

	var start, end string

	if hasStart {
		start = ranges[coreindexes.First]
	}

	if hasEnd {
		end = ranges[coreindexes.Second]
	}

	return &StartEndString{
		BaseRange: &BaseRange{
			RawInput:  raw,
			Separator: sep,
			IsValid:   isValid,
			HasStart:  hasStart,
			HasEnd:    hasEnd,
		},
		Start: start,
		End:   end,
	}
}

// NewStartEndStringUsingLines using first, last index
func NewStartEndStringUsingLines(lines []string) *StartEndString {
	length := len(lines)
	hasStart := length >= 1
	hasEnd := length >= 2
	isValid := length == 2 &&
		hasStart &&
		hasEnd

	var start, end string

	if hasStart {
		start = lines[coreindexes.First]
	}

	if hasEnd {
		end = lines[length-1]
	}

	return &StartEndString{
		BaseRange: &BaseRange{
			RawInput:  constants.EmptyString,
			Separator: constants.EmptyString,
			IsValid:   isValid,
			HasStart:  hasStart,
			HasEnd:    hasEnd,
		},
		Start: start,
		End:   end,
	}
}

func (r *StartEndString) CreateRangeString() *RangeString {
	return &RangeString{
		StartEndString: NewStartEndString(
			r.RawInput,
			r.Separator),
	}
}

func (r *StartEndString) String() string {
	return r.BaseRange.String(r.Start, r.End)
}
