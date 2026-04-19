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

package isany

import (
	"bytes"
	"encoding/json"
)

// JsonEqual
//
//	first checks if string is passed, if yes then only check string.
//	Or, else, marshal and check with error equal if both equal then true.
func JsonEqual(
	left, right any,
) bool {
	leftString, isLeftString := left.(string)
	rightString, isRightString := right.(string)

	if isLeftString && isRightString {
		return leftString == rightString
	}

	leftInteger, isLeftInteger := left.(int)
	rightInteger, isRightInteger := right.(int)

	if isLeftInteger && isRightInteger {
		return leftInteger == rightInteger
	}

	leftBytes, leftErr := json.Marshal(left)
	rightBytes, rightErr := json.Marshal(right)

	if leftErr != nil && rightErr != nil && rightErr.Error() != leftErr.Error() {
		return false
	}

	if leftErr != nil || rightErr != nil {
		return false
	}

	return bytes.Equal(leftBytes, rightBytes)
}
