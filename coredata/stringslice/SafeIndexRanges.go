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

package stringslice

import "github.com/alimtvnetwork/core/constants"

// SafeIndexRanges get safe index range values
// If index range is out of range or not found then empty string will be added to response
//
// Reference Example :
// https://play.golang.org/p/GifVBDSqTJ2
func SafeIndexRanges(
	slice []string,
	startIndexIncluding, endIndexIncluding int,
) []string {
	lastIndex := len(slice) - 1
	requestLength := endIndexIncluding - startIndexIncluding + 1 // +1 for endingIndex or same one if zero

	if requestLength < 0 {
		return []string{}
	}

	responseSlice := MakeLen(requestLength)
	if lastIndex < constants.Zero || requestLength <= 0 {
		return responseSlice
	}

	index := -1

	for i := startIndexIncluding; i <= endIndexIncluding; i++ {
		index++

		if i > lastIndex || i < 0 {
			continue
		}

		responseSlice[index] = slice[i]
	}

	return responseSlice
}
