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

package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================
// Test: Distinct
// ==========================================

func Test_Collection_Distinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.Distinct(col)

		// Handle mixed ExpectedInput types
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {

		// Act
			actual := args.Map{
				"distinctCount": result.Length(),
			}
			for i := 0; i < result.Length(); i++ {
				actual[fmt.Sprintf("item%d", i)] = result.SafeAt(i)
			}

		// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result.Length()))
		}
	}
}

// ==========================================
// Test: DistinctCount
// ==========================================

func Test_Collection_DistinctCount_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", coredynamic.DistinctCount(col)))
	}
}

// ==========================================
// Test: IsDistinct
// ==========================================

func Test_Collection_IsDistinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsDistinctTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", coredynamic.IsDistinct(col)))
	}
}
