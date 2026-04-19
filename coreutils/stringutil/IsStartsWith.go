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

import (
	"strings"
)

func IsStartsWith(
	content, startsWith string,
	isIgnoreCase bool,
) bool {
	if startsWith == "" {
		return true
	}

	if content == "" {
		return startsWith == ""
	}

	if content == startsWith {
		return true
	}

	basePathLength := len(content)
	startsWithLength := len(startsWith)

	if startsWithLength > basePathLength {
		return false
	}

	if isIgnoreCase &&
		basePathLength == startsWithLength &&
		strings.EqualFold(content, startsWith) {
		return true
	}

	if basePathLength <= startsWithLength {
		return false
	}

	remainingText := content[:startsWithLength]

	isCaseSensitive := !isIgnoreCase

	if isCaseSensitive {
		return startsWith == remainingText
	}

	return strings.EqualFold(startsWith, remainingText)
}
