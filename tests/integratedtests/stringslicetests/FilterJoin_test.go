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

func Test_NonEmptySlice_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptySliceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptySlice(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonEmptySlicePtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptySlicePtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptySlicePtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonWhitespace_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespaceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespace(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonWhitespacePtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespacePtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespacePtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonNullStrings_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonNullStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonNullStrings(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonEmptyStrings_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptyStrings(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonEmptyIf_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)
		isFilter := input["isFilter"].(bool)

		// Act
		result := stringslice.NonEmptyIf(isFilter, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_NonEmptyJoin_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyJoinTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptyJoin(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NonEmptyJoinPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonEmptyJoinPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonEmptyJoinPtr(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NonWhitespaceJoin_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespaceJoinTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespaceJoin(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NonWhitespaceJoinPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcNonWhitespaceJoinPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		separator := input["separator"].(string)
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.NonWhitespaceJoinPtr(slice, separator)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_JoinWith_Verification(t *testing.T) {
	for caseIndex, tc := range srcJoinWithTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		separator := input["separator"].(string)
		items := input["items"].([]string)

		// Act
		result := stringslice.JoinWith(separator, items...)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = result != ""
		}
		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Joins_Verification(t *testing.T) {
	for caseIndex, tc := range srcJoinsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		separator := input["separator"].(string)
		items := input["items"].([]string)

		// Act
		result := stringslice.Joins(separator, items...)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_TrimmedEachWords_Verification(t *testing.T) {
	for caseIndex, tc := range srcTrimmedEachWordsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.TrimmedEachWords(slice)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["isNil"]; has {
			actual["isNil"] = result == nil
		}
		if _, has := expected["length"]; has {
			actual["length"] = len(result)
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TrimmedEachWordsPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcTrimmedEachWordsPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw := input["input"]
		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.TrimmedEachWordsPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_TrimmedEachWordsIf_Verification(t *testing.T) {
	for caseIndex, tc := range srcTrimmedEachWordsIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		slice := input["input"].([]string)
		isTrim := input["isTrim"].(bool)

		// Act
		result := stringslice.TrimmedEachWordsIf(isTrim, slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_SplitContentsByWhitespace_Verification(t *testing.T) {
	for caseIndex, tc := range srcSplitContentsByWhitespaceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		str := input["input"].(string)

		// Act
		result := stringslice.SplitContentsByWhitespace(str)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}

func Test_SplitTrimmedNonEmptyAll_Verification(t *testing.T) {
	for caseIndex, tc := range srcSplitTrimmedNonEmptyAllTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		str := input["input"].(string)
		separator := input["separator"].(string)

		// Act
		result := stringslice.SplitTrimmedNonEmptyAll(str, separator)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{"length": len(result)})
	}
}
