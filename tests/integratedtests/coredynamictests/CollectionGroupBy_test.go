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
	"sort"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
)

// ==========================================
// Test: GroupBy
// ==========================================

func Test_Collection_GroupBy_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionGroupByTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		groups := coredynamic.GroupBy(col, func(s string) string {
			if len(s) == 0 {
				return ""
			}
			return string(s[0])
		})

		// Assert
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actLines := make([]string, 0, len(groups))
			for key, group := range groups {
				actLines = append(actLines, fmt.Sprintf("%s:%d", key, group.Length()))
			}
			sort.Strings(actLines)

			actual := args.Map{}
			for i, line := range actLines {
				actual[fmt.Sprintf("group%c", 'A'+i)] = line
			}

			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			actLines := make([]string, 0, len(groups))
			for key, group := range groups {
				actLines = append(actLines, fmt.Sprintf("%s:%d", key, group.Length()))
			}
			sort.Strings(actLines)

			testCase.ShouldBeEqual(t, caseIndex, actLines...)
		}
	}
}

// ==========================================
// Test: GroupByCount
// ==========================================

func Test_Collection_GroupByCount_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionGroupByCountTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		counts := coredynamic.GroupByCount(col, func(s string) string {
			return s
		})

		// Assert
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actLines := make([]string, 0, len(counts))
			for key, count := range counts {
				actLines = append(actLines, fmt.Sprintf("%s:%d", key, count))
			}
			sort.Strings(actLines)

			actual := args.Map{}
			for i, line := range actLines {
				keys := []string{"blueCount", "greenCount", "redCount"}
				if i < len(keys) {
					actual[keys[i]] = line
				}
			}

			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", len(counts)))
		}
	}
}
