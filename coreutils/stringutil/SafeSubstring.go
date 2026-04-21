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

package stringutil

import "github.com/alimtvnetwork/core-v8/constants"

// SafeSubstring content[startAt:endAt]
func SafeSubstring(
	content string,
	startAt, endingLength int,
) string {
	if startAt == -1 && endingLength == -1 {
		return content
	}

	length := len(content)

	if length == 0 {
		return content
	}

	if startAt == -1 {
		return SafeSubstringEnds(content, endingLength)
	}

	if endingLength == -1 {
		return SafeSubstringStarts(content, startAt)
	}

	if startAt < length && endingLength <= length && startAt <= endingLength {
		return content[startAt:endingLength]
	}

	return constants.EmptyString
}
