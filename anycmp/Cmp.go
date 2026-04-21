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

package anycmp

import (
	"github.com/alimtvnetwork/core-v8/corecomparator"
	"github.com/alimtvnetwork/core-v8/isany"
)

// Cmp
//
//	Usages quick determination to say Equal, NotEqual or Inconclusive.
//	Don't use any deep reflect or bytes comparison or string comparison.
//
// Steps:
//   - Returns Equal if both are same pointer
//   - Returns Equal if both are nil
//   - Returns NotEqual if both one nil and another is not
//   - Else, returns Inconclusive.
func Cmp(left, right any) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	isLeftNull, isRightNull := isany.NullLeftRight(left, right)
	isBothEqual := isLeftNull == isRightNull
	isBothDifferent := !isBothEqual

	if isLeftNull && isBothEqual {
		// both null
		return corecomparator.Equal
	} else if isBothDifferent && (isLeftNull || isRightNull) {
		// any null but the other is not
		return corecomparator.NotEqual
	}

	return corecomparator.Inconclusive
}
