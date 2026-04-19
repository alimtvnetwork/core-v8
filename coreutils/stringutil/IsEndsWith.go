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

import "strings"

func IsEndsWith(
	baseStr,
	endsWith string,
	isIgnoreCase bool,
) bool {
	if endsWith == "" {
		return true
	}

	if baseStr == "" {
		return endsWith == ""
	}

	if baseStr == endsWith {
		return true
	}

	basePathLength := len(baseStr)
	endsWithLength := len(endsWith)

	if endsWithLength > basePathLength {
		return false
	}

	if isIgnoreCase &&
		basePathLength == endsWithLength &&
		strings.EqualFold(baseStr, endsWith) {
		return true
	}

	remainingLength := basePathLength - endsWithLength

	if remainingLength < 0 {
		return false
	}

	remainingText := baseStr[remainingLength:]

	isCaseSensitive := !isIgnoreCase

	if isCaseSensitive {
		return endsWith == remainingText
	}

	return strings.EqualFold(endsWith, remainingText)
}
