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

package coreappend

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func PrependAppendAnyItemsToStringsSkipOnNil(
	prependItem, appendItem any,
	anyItems ...any,
) []string {
	slice := make([]string, 0, len(anyItems)+3)

	if prependItem != nil {
		slice = append(
			slice,
			fmt.Sprintf(constants.SprintValueFormat, prependItem))
	}

	for _, item := range anyItems {
		if item == nil {
			continue
		}

		slice = append(
			slice,
			fmt.Sprintf(constants.SprintValueFormat, item))
	}

	if appendItem != nil {
		slice = append(
			slice,
			fmt.Sprintf(constants.SprintValueFormat, appendItem))
	}

	return slice
}
