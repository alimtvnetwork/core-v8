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

package corecmp

import "github.com/alimtvnetwork/core-v8/corecomparator"

func VersionSliceByte(leftVersions, rightVersions []byte) corecomparator.Compare {
	if leftVersions == nil && rightVersions == nil {
		return corecomparator.Equal
	}

	if leftVersions == nil || rightVersions == nil {
		return corecomparator.NotEqual
	}

	leftLen := len(leftVersions)
	rightLen := len(rightVersions)
	minLength := corecomparator.MinLength(
		leftLen,
		rightLen)

	for i := 0; i < minLength; i++ {
		cLeft := leftVersions[i]
		cRight := rightVersions[i]

		if cLeft == cRight {
			continue
		} else if cLeft < cRight {
			return corecomparator.LeftLess
		} else if cLeft > cRight {
			return corecomparator.LeftGreater
		}
	}

	if leftLen == rightLen {
		return corecomparator.Equal
	} else if leftLen < rightLen {
		return corecomparator.LeftLess
	} else if leftLen > rightLen {
		return corecomparator.LeftGreater
	}

	return corecomparator.NotEqual
}
