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

package simplewrap

import "fmt"

// ConditionalWrapWith
//
// Combine start and end char only if missing.
//
// If start missing then add start, if end missing then add end.
func ConditionalWrapWith(
	startChar byte,
	input string,
	endChar byte,
) string {
	length := len(input)

	if length == 0 {
		return fmt.Sprintf(
			"%c%c",
			startChar,
			endChar)
	}

	// more than 2 char
	firstChar := input[0]
	lastChar := input[length-1]
	isPresent :=
		firstChar == startChar &&
			lastChar == endChar

	if isPresent && length == 1 {
		// first one present and single char

		return fmt.Sprintf(
			"%s%c",
			input,
			endChar)
	}

	if isPresent {
		return input
	}

	isRightMissing :=
		firstChar == startChar &&
			lastChar != endChar

	if isRightMissing {
		return fmt.Sprintf(
			"%s%c",
			input,
			endChar)
	}

	isLeftMissing :=
		firstChar != startChar &&
			lastChar == endChar
	if isLeftMissing {
		return fmt.Sprintf(
			"%c%s",
			startChar,
			input,
		)
	}

	// both missing

	return fmt.Sprintf(
		"%c%s%c",
		startChar,
		input,
		endChar,
	)
}
