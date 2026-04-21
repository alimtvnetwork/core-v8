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

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
)

// ==========================================
// Test: Contains
// ==========================================

func Test_Collection_Contains_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionContainsTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.Contains(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IndexOf
// ==========================================

func Test_Collection_IndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIndexOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.IndexOf(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: HasAll
// ==========================================

func Test_Collection_HasAll_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionHasAllTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsStrings("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.HasAll(col, search...)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LastIndexOf
// ==========================================

func Test_Collection_LastIndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionLastIndexOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.LastIndexOf(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Count
// ==========================================

func Test_Collection_Count_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		search, isValid := input.GetAsString("search")

		if !isValid {
			errcore.HandleErrMessage("GetAsString 'search' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.Count(col, search)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
