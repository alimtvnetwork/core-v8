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
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_First_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)

		// Act
		result := stringslice.First(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_FirstOrDefault_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstOrDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.FirstOrDefault(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Last_Verification(t *testing.T) {
	for caseIndex, tc := range srcLastTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)

		// Act
		result := stringslice.Last(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_LastOrDefault_Verification(t *testing.T) {
	for caseIndex, tc := range srcLastOrDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.LastOrDefault(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_IndexAt_Verification(t *testing.T) {
	for caseIndex, tc := range srcIndexAtTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)
		index := input["index"].(int)

		// Act
		result := stringslice.IndexAt(slice, index)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_SafeIndexAt_Verification(t *testing.T) {
	for caseIndex, tc := range srcSafeIndexAtTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		index := input["index"].(int)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.SafeIndexAt(slice, index)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_FirstPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)

		// Act
		result := stringslice.FirstPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_FirstOrDefaultPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstOrDefaultPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.FirstOrDefaultPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_LastPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcLastPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)

		// Act
		result := stringslice.LastPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_LastOrDefaultPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcLastOrDefaultPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.LastOrDefaultPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_FirstOrDefaultWith_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstOrDefaultWithTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		defaultVal := input["default"].(string)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result, ok := stringslice.FirstOrDefaultWith(slice, defaultVal)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
			"ok":     ok,
		})
	}
}

func Test_SafeIndexAtUsingLastIndex_Verification(t *testing.T) {
	for caseIndex, tc := range srcSafeIndexAtUsingLastIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		index := input["index"].(int)
		lastIndex := input["lastIndex"].(int)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.SafeIndexAtUsingLastIndex(slice, index, lastIndex)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_SafeIndexAtWith_Verification(t *testing.T) {
	for caseIndex, tc := range srcSafeIndexAtWithTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		index := input["index"].(int)
		defaultVal := input["default"].(string)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.SafeIndexAtWith(slice, index, defaultVal)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_SafeIndexAtWithPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcSafeIndexAtWithPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		index := input["index"].(int)
		defaultVal := input["default"].(string)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.SafeIndexAtWithPtr(slice, index, defaultVal)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_LastIndexPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcLastIndexPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)

		// Act
		result := stringslice.LastIndexPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_LastSafeIndexPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcLastSafeIndexPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)

		// Act
		result := stringslice.LastSafeIndexPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_IndexesDefault_Verification(t *testing.T) {
	for caseIndex, tc := range srcIndexesDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		indexes := input["indexes"].([]int)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.IndexesDefault(slice, indexes...)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_SafeRangeItems_Verification(t *testing.T) {
	for caseIndex, tc := range srcSafeRangeItemsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		start := input["start"].(int)
		end := input["end"].(int)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.SafeRangeItems(slice, start, end)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_SafeRangeItemsPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcSafeRangeItemsPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		start := input["start"].(int)
		end := input["end"].(int)

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.SafeRangeItemsPtr(slice, start, end)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_FirstLastDefault_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstLastDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		first, last := stringslice.FirstLastDefault(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"first": first,
			"last":  last,
		})
	}
}

func Test_FirstLastDefaultPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstLastDefaultPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		first, last := stringslice.FirstLastDefaultPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"first": first,
			"last":  last,
		})
	}
}

func Test_FirstLastDefaultStatus_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstLastDefaultStatusTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		s := stringslice.FirstLastDefaultStatus(slice)

		// Assert
		actual := args.Map{
			"isValid": s.IsValid,
		}
		if _, has := tc.ExpectedInput.(args.Map)["hasFirst"]; has {
			actual["hasFirst"] = s.HasFirst
		}
		if _, has := tc.ExpectedInput.(args.Map)["hasLast"]; has {
			actual["hasLast"] = s.HasLast
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FirstLastDefaultStatusPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcFirstLastDefaultStatusPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		s := stringslice.FirstLastDefaultStatusPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isValid": s.IsValid,
		})
	}
}
