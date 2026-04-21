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

import "github.com/alimtvnetwork/core-v8/constants"

func InPlaceReverse(list *[]string) *[]string {
	if list == nil {
		return &[]string{}
	}

	nonPtrList := *list
	length := len(nonPtrList)

	if length <= 1 {
		return list
	}

	if length == constants.Capacity2 {
		nonPtrList[0], nonPtrList[1] =
			nonPtrList[1], nonPtrList[0]

		return &nonPtrList
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		nonPtrList[i], nonPtrList[lastIndex-i] =
			nonPtrList[lastIndex-i], nonPtrList[i]
	}

	return &nonPtrList
}
