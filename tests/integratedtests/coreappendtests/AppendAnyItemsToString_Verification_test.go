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

package coreappendtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coreappend"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_AppendAnyItemsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range appendAnyItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")

		// Act
		result := coreappend.AppendAnyItemsToStringSkipOnNil(
			joiner,
			"suffix",
			"item1", "item2",
		)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAnyItemsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAnyItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")

		// Act
		result := coreappend.PrependAnyItemsToStringSkipOnNil(
			joiner,
			"prefix",
			"item1", "item2",
		)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendAnyItemsToString_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")

		// Act
		result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
			joiner,
			"prefix",
			"suffix",
			"item1", "item2",
		)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendSkipNil_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendSkipNilTestCases {
		// Arrange & Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			"prefix",
			"suffix",
			"item1", nil, "item3",
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendNilPrepend_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendNilPrependTestCases {
		// Arrange & Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			nil,
			"suffix",
			"item1",
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PrependAppendNilAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendNilAppendTestCases {
		// Arrange & Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			"prefix",
			nil,
			"item1",
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAppendTestCases {
		// Arrange
		mainMap := map[string]string{"key1": "val1"}
		appendMap := map[string]any{"key2": "val2", "key3": 42}

		// Act
		result := coreappend.MapStringStringAppendMapStringToAnyItems(
			false,
			mainMap,
			appendMap,
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringAppend_SkipEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAppendSkipEmptyTestCases {
		// Arrange
		mainMap := map[string]string{"key1": "val1"}
		appendMap := map[string]any{"key2": "val2", "empty": ""}

		// Act
		result := coreappend.MapStringStringAppendMapStringToAnyItems(
			true,
			mainMap,
			appendMap,
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringAppend_EmptyAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAppendEmptyTestCases {
		// Arrange
		mainMap := map[string]string{"key1": "val1"}
		appendMap := map[string]any{}

		// Act
		result := coreappend.MapStringStringAppendMapStringToAnyItems(
			false,
			mainMap,
			appendMap,
		)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
