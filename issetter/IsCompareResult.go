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

package issetter

import (
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/internal/messages"
)

// IsCompareResult Here left is v, and right is `n`
func (it Value) IsCompareResult(n byte, compare corecomparator.Compare) bool {
	switch compare {
	case corecomparator.Equal:
		return it.IsEqual(n)
	case corecomparator.LeftGreater:
		return it.IsGreater(n)
	case corecomparator.LeftGreaterEqual:
		return it.IsGreaterEqual(n)
	case corecomparator.LeftLess:
		return it.IsLess(n)
	case corecomparator.LeftLessEqual:
		return it.IsLessEqual(n)
	case corecomparator.NotEqual:
		return !it.IsEqual(n)
	default:
		panic(messages.ComparatorOutOfRangeMessage)
	}
}
