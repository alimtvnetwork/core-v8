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

package stringslicetests

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core-v8/coredata/stringslice"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// =============================================================================
// Tests: stringslice.CloneIf
// =============================================================================

func Test_StringSlice_CloneIf(t *testing.T) {
	for caseIndex, testCase := range cloneIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isCloneRaw, _ := input.Get("isClone")
		isClone := isCloneRaw.(bool)
		additionalCapRaw, _ := input.Get("additionalCap")
		additionalCap := additionalCapRaw.(int)
		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)

		var inputSlice []string
		if !isNil {
			inputRaw, _ := input.Get("input")
			inputSlice = inputRaw.([]string)
		}

		// Act
		result := stringslice.CloneIf(isClone, additionalCap, inputSlice)

		actual := args.Map{
			"resultLength": fmt.Sprintf("%d", len(result)),
		}
		for i, v := range result {
			actual[fmt.Sprintf("item%d", i)] = v
		}

		// Check independence
		isIndependentCopy := false
		if len(inputSlice) > 0 && len(result) > 0 {
			isSamePointer := unsafe.Pointer(&inputSlice[0]) == unsafe.Pointer(&result[0])
			isIndependentCopy = !isSamePointer
		} else if isClone {
			isIndependentCopy = true
		}
		actual["isIndependentCopy"] = fmt.Sprintf("%v", isIndependentCopy)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: stringslice.AnyItemsCloneIf
// =============================================================================

func Test_StringSlice_AnyItemsCloneIf(t *testing.T) {
	for caseIndex, testCase := range anyItemsCloneIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isCloneRaw, _ := input.Get("isClone")
		isClone := isCloneRaw.(bool)
		additionalCapRaw, _ := input.Get("additionalCap")
		additionalCap := additionalCapRaw.(int)
		inputRaw, _ := input.Get("input")
		inputSlice := inputRaw.([]any)

		// Act
		result := stringslice.AnyItemsCloneIf(isClone, additionalCap, inputSlice)

		actual := args.Map{
			"resultLength": fmt.Sprintf("%d", len(result)),
		}
		for i, v := range result {
			actual[fmt.Sprintf("item%d", i)] = fmt.Sprintf("%v", v)
		}

		// Check independence
		isIndependentCopy := false
		if len(inputSlice) > 0 && len(result) > 0 {
			isSamePointer := unsafe.Pointer(&inputSlice[0]) == unsafe.Pointer(&result[0])
			isIndependentCopy = !isSamePointer
		} else if isClone {
			isIndependentCopy = true
		}
		actual["isIndependentCopy"] = fmt.Sprintf("%v", isIndependentCopy)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
